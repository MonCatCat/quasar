use crate::bond::{batch_bond, create_share, Bond};
use crate::error::{ContractError, Never, Trap};

use crate::helpers::{
    ack_submsg, create_callback_submsg, create_ibc_ack_submsg, get_ica_address,
    get_usable_compound_balance, unlock_on_error, vec_to_string, IbcMsgKind, IcaMessages,
};
use crate::ibc_lock::Lock;
use crate::ibc_util::{
    calculate_share_out_min_amount, consolidate_exit_pool_amount_into_local_denom,
    do_ibc_join_pool_swap_extern_amount_in, do_ibc_lock_tokens, parse_join_pool,
};
use crate::icq::calc_total_balance;
use crate::start_unbond::{batch_start_unbond, handle_start_unbond_ack};
use crate::state::{
    LpCache, OngoingDeposit, PendingBond, BOND_QUEUE, CHANNELS, CONFIG, IBC_LOCK, IBC_TIMEOUT_TIME, ICA_CHANNEL, ICQ_CHANNEL,
    LP_SHARES, OSMO_LOCK, PENDING_ACK, RECOVERY_ACK, REJOIN_QUEUE, SIMULATED_EXIT_RESULT, SIMULATED_JOIN_RESULT,
    TIMED_OUT, TOTAL_VAULT_BALANCE, TRAPS,
};
use crate::unbond::{batch_unbond, transfer_batch_unbond, PendingReturningUnbonds};
use cosmos_sdk_proto::cosmos::bank::v1beta1::QueryBalanceResponse;

#[cfg(not(feature = "library"))]
use cosmwasm_std::entry_point;

use osmosis_std::types::cosmos::base::v1beta1::Coin as OsmoCoin;
use osmosis_std::types::osmosis::gamm::v1beta1::{
    MsgExitSwapShareAmountInResponse, MsgJoinSwapExternAmountInResponse,
    QueryCalcExitPoolCoinsFromSharesResponse, QueryCalcJoinPoolSharesResponse,
};
use std::str::FromStr;
use std::vec;

use osmosis_std::types::osmosis::gamm::v2::QuerySpotPriceResponse;
use osmosis_std::types::osmosis::lockup::{LockedResponse, MsgLockTokensResponse};
use prost::Message;
use quasar_types::callback::{BondResponse, Callback};
use quasar_types::error::Error as QError;
use quasar_types::ibc::{enforce_order_and_version, ChannelInfo, ChannelType, HandshakeState};
use quasar_types::ica::handshake::enforce_ica_order_and_metadata;
use quasar_types::ica::packet::{ica_send, AckBody};
use quasar_types::ica::traits::Unpack;
use quasar_types::icq::{CosmosResponse, InterchainQueryPacketAck, ICQ_ORDERING};
use quasar_types::{ibc, ica::handshake::IcaMetadata, icq::ICQ_VERSION};

use cosmwasm_std::{
    attr, from_binary, to_binary, Attribute, Binary, Coin, CosmosMsg, Decimal, DepsMut, Env,
    IbcBasicResponse, IbcChannel, IbcChannelCloseMsg, IbcChannelConnectMsg, IbcChannelOpenMsg,
    IbcPacketAckMsg, IbcPacketReceiveMsg, IbcPacketTimeoutMsg, IbcReceiveResponse, IbcTimeout,
    QuerierWrapper, Response, StdError, StdResult, Storage, SubMsg, Uint128, WasmMsg,
};

/// enforces ordering and versioning constraints, this combines ChanOpenInit and ChanOpenTry
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn ibc_channel_open(
    deps: DepsMut,
    _env: Env,
    msg: IbcChannelOpenMsg,
) -> Result<(), ContractError> {
    // save the channel as an channel in ChanOpenInit, we support inits from icq and ica channels
    if msg.channel().version == ICQ_VERSION {
        handle_icq_channel(deps, msg.channel().clone())?;
    } else {
        handle_ica_channel(deps, msg)?;
    }
    Ok(())
}

fn handle_icq_channel(deps: DepsMut, channel: IbcChannel) -> Result<(), ContractError> {
    ibc::enforce_order_and_version(&channel, None, &channel.version, channel.order.clone())?;

    // check the connection id vs the expected connection id
    let config = CONFIG.load(deps.storage)?;
    if config.expected_connection != channel.connection_id {
        return Err(ContractError::IncorrectConnection);
    }

    // save the channel state here
    let info = ChannelInfo {
        id: channel.endpoint.channel_id.clone(),
        counterparty_endpoint: channel.counterparty_endpoint,
        connection_id: channel.connection_id,
        channel_type: ChannelType::Icq {
            channel_ty: channel.version,
        },
        handshake_state: HandshakeState::Init,
    };

    CHANNELS.save(deps.storage, channel.endpoint.channel_id, &info)?;
    Ok(())
}

fn handle_ica_channel(deps: DepsMut, msg: IbcChannelOpenMsg) -> Result<(), ContractError> {
    let channel = msg.channel().clone();
    let metadata: IcaMetadata = serde_json_wasm::from_str(&channel.version).map_err(|error| {
        QError::InvalidIcaMetadata {
            raw_metadata: channel.version.clone(),
            error: error.to_string(),
        }
    })?;

    enforce_ica_order_and_metadata(&channel, None, &metadata)?;

    // compare the expected connection id to the channel connection-id and the ica metadata connection-id
    let config = CONFIG.load(deps.storage)?;
    if &config.expected_connection
        != metadata
            .controller_connection_id()
            .as_ref()
            .ok_or(ContractError::NoConnectionFound)?
    {
        return Err(ContractError::IncorrectConnection);
    }
    if config.expected_connection != channel.connection_id {
        return Err(ContractError::IncorrectConnection);
    }
    // validate that the message is an OpenInit message and not an OpenTry, such that we don't pollute the channel map
    // if let IbcChannelOpenMsg::OpenInit(s) = msg {
    //     return Err(ContractError::InvalidOrder);
    // }
    match msg {
        IbcChannelOpenMsg::OpenInit { channel } => {
            // save the current state of the initializing channel
            let info = ChannelInfo {
                id: channel.endpoint.channel_id.clone(),
                counterparty_endpoint: channel.counterparty_endpoint,
                connection_id: channel.connection_id,
                channel_type: ChannelType::Ica {
                    channel_ty: metadata,
                    counter_party_address: None,
                },
                handshake_state: HandshakeState::Init,
            };
            CHANNELS.save(deps.storage, channel.endpoint.channel_id, &info)?;
            Ok(())
        }
        IbcChannelOpenMsg::OpenTry {
            channel: _,
            counterparty_version: _,
        } => Err(ContractError::IncorrectChannelOpenType),
    }
}

/// record the channel in CHANNEL_INFO, this combines the ChanOpenAck and ChanOpenConfirm steps
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn ibc_channel_connect(
    deps: DepsMut,
    _env: Env,
    msg: IbcChannelConnectMsg,
) -> Result<IbcBasicResponse, ContractError> {
    // try to fetch the connecting channel, we should error if it does not exist\
    let mut info: ChannelInfo = CHANNELS
        .load(deps.storage, msg.channel().endpoint.channel_id.clone())
        .map_err(|err| StdError::GenericErr {
            msg: err.to_string(),
        })?;
    // we need to check the counter party version in try and ack (sometimes here)
    // TODO we can wrap this match in a function in our ibc package

    // TODO think of a better datastructure so we dont have to parse ICA channels like this
    match info.channel_type {
        ChannelType::Icq { ref channel_ty } => {
            enforce_order_and_version(
                msg.channel(),
                msg.counterparty_version(),
                channel_ty.as_str(),
                ICQ_ORDERING,
            )?;
            ICQ_CHANNEL.save(deps.storage, &msg.channel().endpoint.channel_id)?;
            // TODO save the updated state of the ICQ channel
        }
        ChannelType::Ica {
            channel_ty,
            counter_party_address: _,
        } => {
            let counter_party_metadata = enforce_ica_order_and_metadata(
                msg.channel(),
                msg.counterparty_version(),
                &channel_ty,
            )?;

            if counter_party_metadata.is_none() {
                return Err(ContractError::QError(QError::NoCounterpartyIcaAddress));
            }
            let counter_party = counter_party_metadata.unwrap();
            // at this point, we expect a counterparty address, if it's none, we have to error
            if counter_party.address().is_none() {
                return Err(ContractError::NoCounterpartyIcaAddress);
            }
            let addr = counter_party.address();
            if addr.is_none() {
                return Err(ContractError::NoCounterpartyIcaAddress);
            }

            // once we have an Open ICA channel, save it under ICA channel,
            // if a channel already exists, and that channel is not timed out reject incoming OPENS
            // if that channel is timed out, we overwrite the previous ICA channel for the new one
            let channel = ICA_CHANNEL.may_load(deps.storage)?;
            // to reject the msg here, ica should not be timed out
            if channel.is_some() && !TIMED_OUT.load(deps.storage)? {
                return Err(ContractError::IcaChannelAlreadySet);
            }

            // set timed out to false
            TIMED_OUT.save(deps.storage, &false)?;

            ICA_CHANNEL.save(deps.storage, &msg.channel().endpoint.channel_id)?;

            info.channel_type = ChannelType::Ica {
                channel_ty,
                counter_party_address: addr,
            };
            CHANNELS.save(deps.storage, info.id.clone(), &info)?
        }
        ChannelType::Ics20 { channel_ty: _ } => unimplemented!(),
    }

    info.handshake_state = HandshakeState::Open;

    CHANNELS.save(
        deps.storage,
        msg.channel().endpoint.channel_id.clone(),
        &info,
    )?;

    Ok(IbcBasicResponse::default())
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn ibc_channel_close(
    _deps: DepsMut,
    _env: Env,
    channel: IbcChannelCloseMsg,
) -> Result<IbcBasicResponse, ContractError> {
    // TODO look up the channel, in channels, and update the state to closed
    Ok(IbcBasicResponse::new()
        .add_attribute("channel", channel.channel().endpoint.channel_id.clone())
        .add_attribute("connection", channel.channel().connection_id.clone()))
}

/// The lp-strategy cannot receive any packets
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn ibc_packet_receive(
    _deps: DepsMut,
    _env: Env,
    _msg: IbcPacketReceiveMsg,
) -> Result<IbcReceiveResponse, Never> {
    // Contract does not handle packets/queries.
    unimplemented!();
}

#[cfg_attr(not(feature = "library"), entry_point)]
pub fn ibc_packet_ack(
    deps: DepsMut,
    env: Env,
    msg: IbcPacketAckMsg,
) -> Result<IbcBasicResponse, ContractError> {
    // We save the ack binary here for error recovery in case of an join pool recovery
    // this should be cleaned up from state in the ack submsg Ok case
    RECOVERY_ACK.save(
        deps.storage,
        (
            msg.original_packet.sequence,
            msg.original_packet.src.channel_id.clone(),
        ),
        &msg.acknowledgement,
    )?;
    let chan = msg.original_packet.src.channel_id.clone();
    Ok(IbcBasicResponse::new().add_submessage(ack_submsg(deps.storage, env, msg, chan)?))
}

pub fn handle_succesful_ack(
    deps: DepsMut,
    env: Env,
    pkt: IbcPacketAckMsg,
    ack_bin: Binary,
) -> Result<Response, ContractError> {
    let seq = pkt.original_packet.sequence;
    let channel = pkt.original_packet.src.channel_id.clone();

    let kind = PENDING_ACK.load(deps.storage, (seq, channel.clone()))?;
    match kind {
        // a transfer ack means we have sent funds to the ica address, return transfers are handled by the ICA ack
        IbcMsgKind::Transfer { pending, amount } => {
            handle_transfer_ack(deps.storage, env, ack_bin, &pkt, pending, amount)
        }
        IbcMsgKind::Ica(ica_kind) => handle_ica_ack(
            deps.storage,
            deps.querier,
            env,
            ack_bin,
            &pkt,
            ica_kind,
            &seq,
            &channel,
        ),
        IbcMsgKind::Icq => handle_icq_ack(deps.storage, env, ack_bin, &seq, &channel),
    }
}

pub fn handle_transfer_ack(
    storage: &mut dyn Storage,
    env: Env,
    _ack_bin: Binary,
    _pkt: &IbcPacketAckMsg,
    mut pending: PendingBond,
    transferred_amount: Uint128,
) -> Result<Response, ContractError> {
    // once the ibc transfer to the ICA account has succeeded, we send the join pool message
    // we need to save and fetch
    let config = CONFIG.load(storage)?;

    let share_out_min_amount = calculate_share_out_min_amount(storage)?;

    let failed_bonds_amount = REJOIN_QUEUE.iter(storage)?.try_fold(
        Uint128::zero(),
        |acc, val| -> Result<Uint128, ContractError> {
            match val?.raw_amount {
                crate::state::RawAmount::LocalDenom(amount) => Ok(amount + acc),
                crate::state::RawAmount::LpShares(_) => Err(ContractError::IncorrectRawAmount),
            }
        },
    )?;
    let total_amount = transferred_amount + failed_bonds_amount;

    let pending_rejoins: StdResult<Vec<OngoingDeposit>> = REJOIN_QUEUE.iter(storage)?.collect();
    pending.bonds.append(&mut pending_rejoins?);

    let msg = do_ibc_join_pool_swap_extern_amount_in(
        storage,
        env,
        config.pool_id,
        config.base_denom.clone(),
        total_amount,
        share_out_min_amount,
        pending.bonds,
    )?;

    // TODO move this update to after the lock
    TOTAL_VAULT_BALANCE.update(storage, |old| -> Result<Uint128, ContractError> {
        Ok(old.checked_add(total_amount)?)
    })?;

    Ok(Response::new().add_submessage(msg).add_attribute(
        "transfer-ack",
        format!("{}-{}", &total_amount, config.base_denom),
    ))
}

// TODO move the parsing of the ICQ to it's own function, ideally we'd have a type that is contstructed in create ICQ and is parsed from a proto here
pub fn handle_icq_ack(
    storage: &mut dyn Storage,
    env: Env,
    ack_bin: Binary,
    seq: &u64,
    channel: &String,
) -> Result<Response, ContractError> {
    // todo: query flows should be separated by which flowType we're doing (bond, unbond, startunbond)

    let ack: InterchainQueryPacketAck = from_binary(&ack_bin)?;
    let resp: CosmosResponse = CosmosResponse::decode(ack.data.0.as_ref())?;

    // we have only dispatched on query and a single kind at this point
    let raw_balance = QueryBalanceResponse::decode(resp.responses[0].value.as_ref())?
        .balance
        .ok_or(ContractError::BaseDenomNotFound)?
        .amount;

    let base_balance =
        Uint128::new(
            raw_balance
                .parse::<u128>()
                .map_err(|err| ContractError::ParseIntError {
                    error: format!("base_balance:{err}"),
                    value: raw_balance.to_string(),
                })?,
        );

    let usable_base_balance = get_usable_compound_balance(storage, base_balance)?;

    // TODO the quote balance should be able to be compounded aswell
    let _quote_balance = QueryBalanceResponse::decode(resp.responses[1].value.as_ref())?
        .balance
        .ok_or(ContractError::BaseDenomNotFound)?
        .amount;
    // TODO we can make the LP_SHARES cache less error prone here by using the actual state of lp shares
    //  We then need to query locked shares aswell, since they are not part of balance
    let _lp_balance = QueryBalanceResponse::decode(resp.responses[2].value.as_ref())?
        .balance
        .ok_or(ContractError::BaseDenomNotFound)?
        .amount;

    let join_pool = QueryCalcJoinPoolSharesResponse::decode(resp.responses[3].value.as_ref())?;

    let exit_pool =
        QueryCalcExitPoolCoinsFromSharesResponse::decode(resp.responses[4].value.as_ref())?;

    let spot_price = QuerySpotPriceResponse::decode(resp.responses[5].value.as_ref())?.spot_price;
    let lock = LockedResponse::decode(resp.responses[6].value.as_ref())?.lock;
    // parse the locked lp shares on Osmosis, a bit messy
    let gamms = if let Some(lock) = lock {
        lock.coins
    } else {
        vec![]
    };
    let config = CONFIG.load(storage)?;
    let locked_lp_shares = gamms
        .into_iter()
        .find(|val| val.denom == config.pool_denom)
        .unwrap_or(OsmoCoin {
            denom: config.pool_denom.clone(),
            amount: Uint128::zero().to_string(),
        });

    let old_lp_shares = LP_SHARES.load(storage)?;
    // update the locked shares in our cache
    LP_SHARES.update(storage, |mut cache| -> Result<LpCache, ContractError> {
        cache.locked_shares = locked_lp_shares.amount.parse()?;
        Ok(cache)
    })?;

    let spot_price =
        Decimal::from_str(spot_price.as_str()).map_err(|err| ContractError::ParseDecError {
            error: err,
            value: spot_price,
        })?;

    let total_balance = calc_total_balance(
        storage,
        usable_base_balance,
        &exit_pool.tokens_out,
        spot_price,
    )?;

    let parsed_exit_pool_out =
        consolidate_exit_pool_amount_into_local_denom(storage, &exit_pool.tokens_out, spot_price)?;

    TOTAL_VAULT_BALANCE.save(storage, &total_balance)?;

    let parsed_join_pool_out = parse_join_pool(storage, join_pool)?;

    SIMULATED_JOIN_RESULT.save(storage, &parsed_join_pool_out)?;
    SIMULATED_EXIT_RESULT.save(storage, &parsed_exit_pool_out)?;

    // get bond information before batching bonds
    let mut bonds: Vec<Bond> = vec![];
    if !BOND_QUEUE.is_empty(storage)? {
        bonds = BOND_QUEUE
            .iter(storage)?
            .collect::<StdResult<Vec<Bond>>>()?;
    }

    // todo move this to below into the lock decisions
    let bond = batch_bond(storage, &env, total_balance)?;

    let mut msges = Vec::new();
    let mut attrs = Vec::new();
    // if queues had items, msges should be some, so we add the ibc submessage, if there were no items in a queue, we don't have a submsg to add
    // if we have a bond, start_unbond or unbond msg, we lock the repsective lock

    // todo rewrite into flat if/else ifs

    if let Some(msg) = bond {
        msges.push(msg);

        let owners = bonds.iter().map(|b| b.owner.to_string()).collect();
        let bond_ids = bonds.iter().map(|b| b.bond_id.clone()).collect();
        let amounts = bonds
            .iter()
            // TODO: is base denom the correct denom here?
            .map(|b| b.amount.to_string() + &config.base_denom)
            .collect();

        // E2: BondPacket
        attrs = vec![
            attr("action", "bond_packet"),
            attr("primitive_address", env.contract.address.to_string()),
            attr("owners", vec_to_string(owners)),
            attr("bond_ids", vec_to_string(bond_ids)),
            attr("amounts", vec_to_string(amounts)),
            attr("packet_sequence", seq.to_string()),
            attr("channel_id", channel.clone()),
            attr("data", ""),
        ];

        IBC_LOCK.update(storage, |lock| -> Result<Lock, ContractError> {
            Ok(lock.lock_bond())
        })?;
    } else {
        attrs.push(attr("bond-status", "empty"));
        if let Some(msg) = batch_start_unbond(storage, &env)? {
            msges.push(msg);
            attrs.push(attr("start-unbond-status", "starting-unbond"));
            IBC_LOCK.update(storage, |lock| -> Result<Lock, ContractError> {
                Ok(lock.lock_start_unbond())
            })?;
        } else {
            attrs.push(attr("start-unbond-status", "empty"));
            if let Some(msg) = batch_unbond(storage, &env, old_lp_shares)? {
                msges.push(msg);
                attrs.push(attr("unbond-status", "unbonding"));
                IBC_LOCK.update(storage, |lock| -> Result<Lock, ContractError> {
                    Ok(lock.lock_unbond())
                })?;
            } else {
                attrs.push(attr("unbond-status", "empty"));
            }
        }
    }

    Ok(Response::new().add_submessages(msges).add_attributes(attrs))
}

pub fn handle_ica_ack(
    storage: &mut dyn Storage,
    querier: QuerierWrapper,
    env: Env,
    ack_bin: Binary,
    _pkt: &IbcPacketAckMsg,
    ica_kind: IcaMessages,
    seq: &u64,
    channel: &String,
) -> Result<Response, ContractError> {
    match ica_kind {
        IcaMessages::JoinSwapExternAmountIn(mut data) => {
            handle_join_pool(storage, &env, ack_bin, &mut data)
        }
        IcaMessages::LockTokens(data, lp_shares) => handle_lock_tokens_ack(
            storage, &env, data, lp_shares, ack_bin, querier, seq, channel,
        ),
        IcaMessages::BeginUnlocking(data, total) => {
            handle_start_unbond_ack(storage, querier, &env, data, total)
        }
        IcaMessages::ExitPool(data) => handle_exit_pool_ack(storage, &env, data, ack_bin),
        // TODO decide where we unlock the transfer ack unlock, here or in the ibc hooks receive
        IcaMessages::ReturnTransfer(data) => handle_return_transfer_ack(storage, querier, data),
        // After a RecoveryExitPool, we do a return transfer that should hit RecoveryReturnTransfer
        IcaMessages::RecoveryExitPool(_pending) => todo!(),
        // After a RecoveryReturnTransfer, we save the funds to a local map, to be claimed by vaults when a users asks
        IcaMessages::RecoveryReturnTransfer(_pending) => todo!(),
    }
}

// fn handle_recovery_return_transfer(
//     storage: &mut dyn Storage,
//     pending: PendingReturningRecovery,

// ) -> Result<Response, ContractError> {
//     // if we have the succesfully received the recovery, we create an entry
//     for p in pending.returning {
//         if let RawAmount::LocalDenom(val) = p.amount {
//             CLAIMABLE_FUNDS.save(storage, (p.owner, p.id), &val)?;
//         } else {
//             return Err(ContractError::IncorrectRawAmount);
//         }
//         // remove the error from TRAPS
//         TRAPS.remove(storage, (pending.trapped_id, ));
//     }
//     todo!()
// }

fn handle_join_pool(
    storage: &mut dyn Storage,
    env: &Env,
    ack_bin: Binary,
    data: &mut PendingBond,
) -> Result<Response, ContractError> {
    // TODO move the below locking logic to a separate function
    // get the ica address of the channel id
    let ica_channel = ICA_CHANNEL.load(storage)?;
    let ica_addr = get_ica_address(storage, ica_channel.clone())?;
    let ack = AckBody::from_bytes(ack_bin.0.as_ref())?.to_any()?;
    let resp = MsgJoinSwapExternAmountInResponse::unpack(ack)?;
    let shares_out = Uint128::new(resp.share_out_amount.parse::<u128>().map_err(|err| {
        ContractError::ParseIntError {
            error: format!("{err}"),
            value: resp.share_out_amount,
        }
    })?);

    let denom = CONFIG.load(storage)?.pool_denom;

    LP_SHARES.update(
        storage,
        |mut old: LpCache| -> Result<LpCache, ContractError> {
            old.d_unlocked_shares = old.d_unlocked_shares.checked_add(shares_out)?;
            Ok(old)
        },
    )?;

    data.update_raw_amount_to_lp(shares_out)?;

    let msg = do_ibc_lock_tokens(
        storage,
        ica_addr,
        vec![Coin {
            denom,
            amount: shares_out,
        }],
    )?;

    let channel = ICA_CHANNEL.load(storage)?;

    let outgoing = ica_send(
        msg,
        ica_channel,
        IbcTimeout::with_timestamp(env.block.time.plus_seconds(IBC_TIMEOUT_TIME)),
    )?;

    let msg = create_ibc_ack_submsg(
        storage,
        IbcMsgKind::Ica(IcaMessages::LockTokens(data.clone(), shares_out)),
        outgoing,
        channel,
    )?;
    Ok(Response::new().add_submessage(msg))
}

fn handle_lock_tokens_ack(
    storage: &mut dyn Storage,
    env: &Env,
    data: PendingBond,
    total_lp_shares: Uint128,
    ack_bin: Binary,
    querier: QuerierWrapper,
    seq: &u64,
    channel: &String,
) -> Result<Response, ContractError> {
    let ack = AckBody::from_bytes(ack_bin.0.as_ref())?.to_any()?;
    let resp = MsgLockTokensResponse::unpack(ack)?;

    // save the lock id in the contract
    OSMO_LOCK.save(storage, &resp.id)?;

    LP_SHARES.update(storage, |mut old| -> Result<LpCache, ContractError> {
        old.d_unlocked_shares =
            old.d_unlocked_shares
                .checked_sub(total_lp_shares)
                .map_err(|err| {
                    ContractError::TracedOverflowError(
                        err,
                        "update_unlocked_deposit_shares".to_string(),
                    )
                })?;
        old.locked_shares = old
            .locked_shares
            .checked_add(total_lp_shares)
            .map_err(|err| {
                ContractError::TracedOverflowError(err, "update_locked_shares".to_string())
            })?;
        Ok(old)
    })?;

    let mut callback_submsgs: Vec<SubMsg> = vec![];
    let mut owners = vec![];
    let mut bond_ids = vec![];

    for claim in data.bonds {
        let share_amount = create_share(storage, &claim.owner, &claim.bond_id, claim.claim_amount)?;

        println!("###DEBUG: {}", claim.owner.as_str());

        if querier
            .query_wasm_contract_info(claim.owner.as_str())
            .is_ok()
        {
            let wasm_msg = WasmMsg::Execute {
                contract_addr: claim.owner.to_string(),
                msg: to_binary(&Callback::BondResponse(BondResponse {
                    share_amount,
                    bond_id: claim.bond_id.clone(),
                }))?,
                funds: vec![],
            };

            // convert wasm_msg into cosmos_msg to be handled in create_callback_submsg
            let cosmos_msg = CosmosMsg::Wasm(wasm_msg);
            callback_submsgs.push(create_callback_submsg(
                storage,
                cosmos_msg,
                claim.owner.clone(),
                claim.bond_id.clone(),
            )?);
        }

        owners.push(claim.owner.to_string());
        bond_ids.push(claim.bond_id);
    }

    // set the bond lock state to unlocked
    IBC_LOCK.update(storage, |old| -> Result<Lock, StdError> {
        Ok(old.unlock_bond())
    })?;

    // TODO, do we want to also check queue state? and see if we can already start a new execution?
    Ok(Response::new()
        .add_submessages(callback_submsgs)
        // E3: BondAcknowledgement
        .add_attributes(vec![
            ("action", "bond_acknowledgment"),
            ("primitive_address", &env.contract.address.to_string()),
            ("owners", &vec_to_string(owners)),
            ("bond_ids", &vec_to_string(bond_ids)),
            ("packet_sequence", &seq.to_string()),
            ("channel_id", channel),
            ("data", ""),
        ]))
}

fn handle_exit_pool_ack(
    storage: &mut dyn Storage,
    env: &Env,
    mut data: PendingReturningUnbonds,
    ack_bin: Binary,
) -> Result<Response, ContractError> {
    let ack = AckBody::from_bytes(ack_bin.0.as_ref())?.to_any()?;
    let msg = MsgExitSwapShareAmountInResponse::unpack(ack)?;
    let total_exited_tokens =
        Uint128::new(msg.token_out_amount.parse::<u128>().map_err(|err| {
            ContractError::ParseIntError {
                error: format!("{err}"),
                value: msg.token_out_amount,
            }
        })?);

    // we don't need the sum of the lp tokens returned by lp_to_local_denom here
    let _ = data.lp_to_local_denom(total_exited_tokens)?;

    let sub_msg = transfer_batch_unbond(storage, env, data, total_exited_tokens)?;
    Ok(Response::new()
        .add_submessage(sub_msg)
        .add_attribute("transfer-funds", total_exited_tokens.to_string()))
}

fn handle_return_transfer_ack(
    storage: &mut dyn Storage,
    _querier: QuerierWrapper,
    _data: PendingReturningUnbonds,
) -> Result<Response, ContractError> {
    IBC_LOCK.update(storage, |lock| -> Result<Lock, ContractError> {
        Ok(lock.unlock_unbond())
    })?;

    Ok(Response::new().add_attribute("return-transfer", "success"))
}

pub fn handle_failing_ack(
    deps: DepsMut,
    _env: Env,
    pkt: IbcPacketAckMsg,
    error: String,
) -> Result<Response, ContractError> {
    // TODO we can expand error handling here to fetch the packet by the ack and add easy retries or something
    let step = PENDING_ACK.load(
        deps.storage,
        (
            pkt.original_packet.sequence,
            pkt.original_packet.src.channel_id.clone(),
        ),
    )?;
    unlock_on_error(deps.storage, &step)?;
    TRAPS.save(
        deps.storage,
        (
            pkt.original_packet.sequence,
            pkt.original_packet.src.channel_id,
        ),
        &Trap {
            error: format!("packet failure: {error}"),
            step: step.clone(),
            last_succesful: false,
        },
    )?;
    Ok(Response::new().add_attribute("ibc-error", error.as_str()))
}

// if an ICA packet is timed out, we need to reject any further packets (only to the ICA channel or in total -> easiest in total until a new ICA channel is created)
// once time out variable is set, a new ICA channel needs to be able to be opened for the contract to function and the ICA channel val and channels map need to be updated
// what do we do with the trapped errors packets, are they able to be recovered over the new ICA channel?
#[cfg_attr(not(feature = "library"), entry_point)]
pub fn ibc_packet_timeout(
    deps: DepsMut,
    _env: Env,
    msg: IbcPacketTimeoutMsg,
) -> Result<IbcBasicResponse, ContractError> {
    on_packet_timeout(
        deps,
        msg.packet.sequence,
        msg.packet.src.channel_id,
        "timeout".to_string(),
        true,
    )
}

pub(crate) fn on_packet_timeout(
    deps: DepsMut,
    sequence: u64,
    channel: String,
    error: String,
    should_unlock: bool,
) -> Result<IbcBasicResponse, ContractError> {
    let step = PENDING_ACK.load(deps.storage, (sequence, channel.clone()))?;
    if should_unlock {
        unlock_on_error(deps.storage, &step)?;
    }
    if let IbcMsgKind::Ica(_) = &step {
        TIMED_OUT.save(deps.storage, &true)?
    }
    TRAPS.save(
        deps.storage,
        (sequence, channel),
        &Trap {
            error: format!("packet failure: {error}"),
            step,
            last_succesful: false,
        },
    )?;
    Ok(IbcBasicResponse::default())
}

#[cfg(test)]
mod tests {

    use cosmwasm_std::{
        testing::{mock_dependencies, mock_env, MockQuerier},
        Addr, ContractInfoResponse, ContractResult, IbcEndpoint, IbcOrder, QuerierResult,
        WasmQuery,
    };

    use crate::{
        state::{Config, OngoingDeposit, RawAmount, BONDING_CLAIMS, SIMULATED_JOIN_AMOUNT_IN},
        test_helpers::default_setup,
    };

    use super::*;

    #[test]
    fn handle_icq_ack_works() {
        let mut deps = mock_dependencies();
        let env = mock_env();
        default_setup(deps.as_mut().storage).unwrap();

        CONFIG
            .save(
                deps.as_mut().storage,
                &Config {
                    lock_period: 100,
                    pool_id: 1,
                    pool_denom: "gamm/pool/1".to_string(),
                    base_denom:
                        "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2"
                            .to_string(),
                    quote_denom: "uosmo".to_string(),
                    local_denom:
                        "ibc/FA0006F056DB6719B8C16C551FC392B62F5729978FC0B125AC9A432DBB2AA1A5"
                            .to_string(),
                    local_denom: "ibc/local_osmo".to_string(),
                    transfer_channel: "channel-0".to_string(),
                    return_source_channel: "channel-0".to_string(),
                    expected_connection: "connection-0".to_string(),
                },
            )
            .unwrap();
        LP_SHARES
            .save(
                deps.as_mut().storage,
                &LpCache {
                    locked_shares: Uint128::new(1000),
                    w_unlocked_shares: Uint128::zero(),
                    d_unlocked_shares: Uint128::zero(),
                },
            )
            .unwrap();
        SIMULATED_JOIN_AMOUNT_IN
            .save(deps.as_mut().storage, &Uint128::zero())
            .unwrap();

        // base64 of '{"data":"Chs6FAoSCgV1b3NtbxIJMTkyODcwODgySNW/pQQKUjpLCkkKRGliYy8yNzM5NEZCMDkyRDJFQ0NENTYxMjNDNzRGMzZFNEMxRjkyNjAwMUNFQURBOUNBOTdFQTYyMkIyNUY0MUU1RUIyEgEwSNW/pQQKGToSChAKC2dhbW0vcG9vbC8xEgEwSNW/pQQKFjoPCgEwEgoKBXVvc21vEgEwSNW/pQQKcTpqClIKRGliYy8yNzM5NEZCMDkyRDJFQ0NENTYxMjNDNzRGMzZFNEMxRjkyNjAwMUNFQURBOUNBOTdFQTYyMkIyNUY0MUU1RUIyEgoxMDg5ODQ5Nzk5ChQKBXVvc21vEgsxNTQyOTM2Mzg2MEjVv6UECh06FgoUMC4wNzA2MzQ3ODUwMDAwMDAwMDBI1b+lBAqMATqEAQqBAQj7u2ISP29zbW8xd212ZXpscHNrNDB6M3pmc3l5ZXgwY2Q4ZHN1bTdnenVweDJxZzRoMHVhdms3dHh3NHNlcXE3MmZrbRoECIrqSSILCICSuMOY/v///wEqJwoLZ2FtbS9wb29sLzESGDEwODE3NDg0NTgwODQ4MDkyOTUyMDU1MUjVv6UE"}'
        let ack_bin = Binary::from_base64("eyJkYXRhIjoiQ2xnNlVRcFBDa1JwWW1Ndk1qY3pPVFJHUWpBNU1rUXlSVU5EUkRVMk1USXpRemMwUmpNMlJUUkRNVVk1TWpZd01ERkRSVUZFUVRsRFFUazNSVUUyTWpKQ01qVkdOREZGTlVWQ01oSUhNalV3TURBd01FaUdydk1FQ2hNNkRBb0tDZ1YxYjNOdGJ4SUJNRWlHcnZNRUNoazZFZ29RQ2d0bllXMXRMM0J2YjJ3dk1SSUJNRWlHcnZNRUNtMDZaZ29VTWpneU5EWTVOVE0yTnpFek9EQXpNems1T1RnU1RncEVhV0pqTHpJM016azBSa0l3T1RKRU1rVkRRMFExTmpFeU0wTTNORVl6TmtVMFF6RkdPVEkyTURBeFEwVkJSRUU1UTBFNU4wVkJOakl5UWpJMVJqUXhSVFZGUWpJU0JqVXdNREF3TUVpR3J2TUVDbkE2YVFwUkNrUnBZbU12TWpjek9UUkdRakE1TWtReVJVTkRSRFUyTVRJelF6YzBSak0yUlRSRE1VWTVNall3TURGRFJVRkVRVGxEUVRrM1JVRTJNakpDTWpWR05ERkZOVVZDTWhJSk5qTTJNell6TVRJMkNoUUtCWFZ2YzIxdkVnc3hNVGd4TXpnek5ESXdNVWlHcnZNRUNoNDZGd29WTVRndU5UWTBOakV4TkRnd01EQXdNREF3TURBd1NJYXU4d1FLaXdFNmd3RUtnQUVJNWZWa0VqOXZjMjF2TVhWNU5XRnlNSGxoYXpseU4zZ3pjMmRuWkRkbmMyNXlhbTEwYTNjeWRqWmxkRGgxZUhKeGF6TnRObkJvTkdaMWNYYzVaSEUyYTJGNmJYTWFCQWlCNmtraUN3aUFrcmpEbVA3Ly8vOEJLaVlLQzJkaGJXMHZjRzl2YkM4eEVoYzNNVGszTXpJMU56QTJOVGcwTnprek56QXpPVFV6TlVpR3J2TUUifQ==").unwrap();
        // queues are empty at this point so we just expect a succesful response without anyhting else
        handle_icq_ack(
            deps.as_mut().storage,
            env,
            ack_bin,
            &1,
            &"channel-25".to_string(),
        )
        .unwrap();
    }

    #[test]
    fn handle_ica_channel_works() {
        let mut deps = mock_dependencies();
        default_setup(deps.as_mut().storage).unwrap();

        let endpoint = IbcEndpoint {
            port_id: "wasm.my_addr".to_string(),
            channel_id: "channel-1".to_string(),
        };
        let counterparty_endpoint = IbcEndpoint {
            port_id: "icahost".to_string(),
            channel_id: "channel-2".to_string(),
        };

        let version = r#"{"version":"ics27-1","encoding":"proto3","tx_type":"sdk_multi_msg","controller_connection_id":"connection-0","host_connection_id":"connection-0"}"#.to_string();
        let channel = IbcChannel::new(
            endpoint,
            counterparty_endpoint.clone(),
            IbcOrder::Ordered,
            version,
            "connection-0".to_string(),
        );

        let msg = IbcChannelOpenMsg::OpenInit {
            channel: channel.clone(),
        };

        handle_ica_channel(deps.as_mut(), msg.clone()).unwrap();

        let expected = ChannelInfo {
            id: channel.endpoint.channel_id.clone(),
            counterparty_endpoint,
            connection_id: "connection-0".to_string(),
            channel_type: ChannelType::Ica {
                channel_ty: IcaMetadata::with_connections(
                    "connection-0".to_string(),
                    "connection-0".to_string(),
                ),
                counter_party_address: None,
            },
            handshake_state: HandshakeState::Init,
        };
        assert_eq!(
            CHANNELS
                .load(deps.as_ref().storage, channel.endpoint.channel_id)
                .unwrap(),
            expected
        )
    }

    #[test]
    fn handle_ica_channel_open_try_errors() {
        let mut deps = mock_dependencies();
        default_setup(deps.as_mut().storage).unwrap();

        let endpoint = IbcEndpoint {
            port_id: "wasm.my_addr".to_string(),
            channel_id: "channel-1".to_string(),
        };
        let counterparty_endpoint = IbcEndpoint {
            port_id: "icahost".to_string(),
            channel_id: "channel-2".to_string(),
        };

        let version = r#"{"version":"ics27-1","encoding":"proto3","tx_type":"sdk_multi_msg","controller_connection_id":"connection-0","host_connection_id":"connection-0"}"#.to_string();
        let channel = IbcChannel::new(
            endpoint,
            counterparty_endpoint.clone(),
            IbcOrder::Ordered,
            version,
            "connection-0".to_string(),
        );

        let msg = IbcChannelOpenMsg::OpenTry {
            channel: channel.clone(),
            counterparty_version: "1".to_string(),
        };

        let err = handle_ica_channel(deps.as_mut(), msg.clone()).unwrap_err();
        assert_eq!(err, ContractError::IncorrectChannelOpenType);
    }

    #[test]
    fn handle_icq_ack_events_bond_queue() {
        let mut deps = mock_dependencies();
        let env = mock_env();
        default_setup(deps.as_mut().storage).unwrap();

        CONFIG
            .save(
                deps.as_mut().storage,
                &Config {
                    lock_period: 100,
                    pool_id: 1,
                    pool_denom: "gamm/pool/1".to_string(),
                    base_denom: "uosmo".to_string(),
                    quote_denom:
                        "ibc/D176154B0C63D1F9C6DCFB4F70349EBF2E2B5A87A05902F57A6AE92B863E9AEC"
                            .to_string(),
                    local_denom: "ibc/local_osmo".to_string(),
                    transfer_channel: "channel-0".to_string(),
                    return_source_channel: "channel-0".to_string(),
                    expected_connection: "connection-0".to_string(),
                },
            )
            .unwrap();

        SIMULATED_JOIN_AMOUNT_IN
            .save(deps.as_mut().storage, &Uint128::zero())
            .unwrap();

        let ack_bin = Binary::from_base64("eyJkYXRhIjoiQ2hzNkZBb1NDZ1YxYjNOdGJ4SUpNVGM1TVRjME5EYzNTTlcvcFFRS1VqcExDa2tLUkdsaVl5OUVNVGMyTVRVMFFqQkROak5FTVVZNVF6WkVRMFpDTkVZM01ETTBPVVZDUmpKRk1rSTFRVGczUVRBMU9UQXlSalUzUVRaQlJUa3lRamcyTTBVNVFVVkRFZ0V3U05XL3BRUUtHem9VQ2hJS0RXZGhiVzB2Y0c5dmJDODRNek1TQVRCSTFiK2xCQW9IQ0JKSTFiK2xCQXB6T213S1V3cEVhV0pqTDBReE56WXhOVFJDTUVNMk0wUXhSamxETmtSRFJrSTBSamN3TXpRNVJVSkdNa1V5UWpWQk9EZEJNRFU1TURKR05UZEJOa0ZGT1RKQ09EWXpSVGxCUlVNU0N6azBNRFl3TWpNMU1UY3hDaFVLQlhWdmMyMXZFZ3d4TWpNNE9EUTJNRGN6TVRCSTFiK2xCQW9kT2hZS0ZEQXVPVEl4TlRrNU9ESXdNREF3TURBd01EQXdTTlcvcFFRS2l3RTZnd0VLZ0FFSS9MdGlFajl2YzIxdk1YQnpjMlo2Y0Roa05tZzFjR3R6Wm5sak5tdzFNamRtYUdkMlpHcGpOVE0zZFhWbmRIQm5NbVUwZDI1M1pIRjFlWFpxWVhGa2MyaHdZV2dhQkFpSzZra2lDd2lBa3JqRG1QNy8vLzhCS2lZS0RXZGhiVzB2Y0c5dmJDODRNek1TRlRFMk1qQXhOVFU0T1RjM01ERXpNems0TURRM01ralZ2NlVFIn0").unwrap();

        BOND_QUEUE
            .push_back(
                &mut deps.storage,
                &Bond {
                    amount: Uint128::one(),
                    owner: Addr::unchecked("vault_1".to_string()),
                    bond_id: "1".to_string(),
                },
            )
            .unwrap();

        BOND_QUEUE
            .push_back(
                &mut deps.storage,
                &Bond {
                    amount: Uint128::new(2),
                    owner: Addr::unchecked("vault_1".to_string()),
                    bond_id: "2".to_string(),
                },
            )
            .unwrap();

        IBC_LOCK.save(deps.as_mut().storage, &Lock::new()).unwrap();

        let res = handle_icq_ack(
            deps.as_mut().storage,
            env,
            ack_bin,
            &100,
            &"channel-25".to_string(),
        )
        .unwrap();
        assert_eq!(res.messages.len(), 1);
        assert_eq!(
            res.attributes,
            vec![
                attr("action", "bond_packet"),
                attr("primitive_address", "cosmos2contract"),
                attr("owners", "['vault_1','vault_1']"),
                attr("bond_ids", "['1','2']"),
                attr("amounts", "['1uosmo','2uosmo']"),
                attr("packet_sequence", 100.to_string()),
                attr("channel_id", "channel-25"),
                attr("data", ""),
            ]
        );
    }

    #[test]
    fn handle_lock_tokens_ack_events_bond_queue() {
        let mut deps = mock_dependencies();
        let env = mock_env();
        default_setup(deps.as_mut().storage).unwrap();

        CONFIG
            .save(
                deps.as_mut().storage,
                &Config {
                    lock_period: 100,
                    pool_id: 1,
                    pool_denom: "gamm/pool/1".to_string(),
                    base_denom: "uosmo".to_string(),
                    quote_denom:
                        "ibc/D176154B0C63D1F9C6DCFB4F70349EBF2E2B5A87A05902F57A6AE92B863E9AEC"
                            .to_string(),
                    local_denom: "ibc/local_osmo".to_string(),
                    transfer_channel: "channel-0".to_string(),
                    return_source_channel: "channel-0".to_string(),
                    expected_connection: "connection-0".to_string(),
                },
            )
            .unwrap();

        IBC_LOCK.save(deps.as_mut().storage, &Lock::new()).unwrap();

        let pending = PendingBond {
            bonds: vec![
                OngoingDeposit {
                    claim_amount: Uint128::new(1),
                    raw_amount: RawAmount::LocalDenom(Uint128::new(1000)),
                    owner: Addr::unchecked("vault_1"),
                    bond_id: "1".to_string(),
                },
                OngoingDeposit {
                    claim_amount: Uint128::new(2),
                    raw_amount: RawAmount::LocalDenom(Uint128::new(999)),
                    owner: Addr::unchecked("vault_1"),
                    bond_id: "2".to_string(),
                },
            ],
        };

        BONDING_CLAIMS
            .save(
                deps.as_mut().storage,
                (&Addr::unchecked("vault_1".to_string()), "1"),
                &Uint128::new(1),
            )
            .unwrap();

        BONDING_CLAIMS
            .save(
                deps.as_mut().storage,
                (&Addr::unchecked("vault_1".to_string()), "2"),
                &Uint128::new(2),
            )
            .unwrap();

        let ack_bin =
            Binary::from_base64("CiUKHS9vc21vc2lzLmxvY2t1cC5Nc2dMb2NrVG9rZW5zEgQI/Lti").unwrap();

        let mut info = ContractInfoResponse::default();
        info.admin = Some("vault_1".to_string());

        let mut q = MockQuerier::default();
        q.update_wasm(move |q: &WasmQuery| -> QuerierResult {
            match q {
                WasmQuery::ContractInfo { contract_addr: _ } => {
                    QuerierResult::Ok(ContractResult::Ok(to_binary(&info.admin).unwrap()))
                }
                _ => unreachable!(),
            }
        });
        let w: QuerierWrapper = QuerierWrapper::new(&deps.querier);

        let test = w.query_wasm_contract_info(&env.contract.address); //pending.bonds[0].owner.as_str());

        println!("###DEBUG: {:?}", test);

        let res = handle_lock_tokens_ack(
            &mut deps.storage,
            &env,
            pending,
            Uint128::new(3),
            ack_bin.clone(),
            w,
            &100,
            &"channel-25".to_string(),
        )
        .unwrap();

        // TODO: check why this is 0 and not 1
        assert_eq!(res.messages.len(), 0);
        assert_eq!(
            res.attributes,
            vec![
                attr("action", "bond_acknowledgment"),
                attr("primitive_address", "cosmos2contract"),
                attr("owners", "['vault_1','vault_1']"),
                attr("bond_ids", "['1','2']"),
                attr("packet_sequence", 100.to_string()),
                attr("channel_id", "channel-25"),
                attr("data", ""),
            ]
        );
    }
}
