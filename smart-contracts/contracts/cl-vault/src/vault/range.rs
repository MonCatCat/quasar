use std::str::FromStr;

use cosmwasm_std::{
    Addr, Decimal, Decimal256, Deps, DepsMut, Env, MessageInfo, QuerierWrapper, Response, Storage,
    SubMsg, Uint128,
};
use cw_utils::nonpayable;
use osmosis_std::types::cosmos::base::v1beta1::Coin as OsmoCoin;

use crate::{
    concentrated_liquidity::get_position,
    helpers::{
        get_deposit_amounts_for_liquidity_needed, get_liquidity_needed_for_tokens, with_slippage,
    },
    math::tick::price_to_tick,
    reply::Replies,
    state::{ADMIN_ADDRESS, POOL_CONFIG, RANGE_ADMIN, VAULT_CONFIG},
    swap::SwapDirection,
    ContractError,
};

pub fn execute_modify_range(
    deps: DepsMut,
    env: Env,
    info: MessageInfo,
    lower_price: Uint128,
    upper_price: Uint128,
) -> Result<Response, ContractError> {
    // let lower_tick = price_to_tick(price, exponent_at_price_one)

    let storage = deps.storage;
    let querier = deps.querier;

    let lower_tick = price_to_tick(storage, Decimal256::from_atomics(lower_price, 0)?)?;
    let upper_tick = price_to_tick(storage, Decimal256::from_atomics(upper_price, 0)?)?;

    execute_modify_range_ticks(
        storage,
        &querier,
        env,
        info,
        lower_price,
        upper_price,
        lower_tick,
        upper_tick,
    )
}

fn assert_range_admin(storage: &mut dyn Storage, sender: &Addr) -> Result<(), ContractError> {
    let admin = RANGE_ADMIN.load(storage)?;
    if &admin != sender {
        return Err(ContractError::Unauthorized {});
    }
    Ok(())
}

fn get_range_admin(deps: Deps) -> Result<Addr, ContractError> {
    Ok(RANGE_ADMIN.load(deps.storage)?)
}

pub fn execute_modify_range_ticks(
    storage: &mut dyn Storage,
    querier: &QuerierWrapper,
    env: Env,
    info: MessageInfo,
    lower_price: Uint128,
    upper_price: Uint128,
    lower_tick: i128,
    upper_tick: i128,
) -> Result<Response, ContractError> {
    assert_range_admin(storage, &info.sender)?;

    let pool_config = POOL_CONFIG.load(storage)?;
    let vault_config = VAULT_CONFIG.load(storage)?;

    // This function is the entrypoint into the dsm routine that will go through the following steps
    // * how much liq do we have in current range
    // * so how much of each asset given liq would we have at current price
    // * how much of each asset do we need to move to get to new range
    // * deposit up to max liq we can right now, then swap remaining over and deposit again

    // this will error if we dont have a position anyway
    let position = get_position(storage, &querier, &env)?;

    let liquidity = match position.position {
        Some(position) => Decimal::from_str(&position.liquidity)
            .expect("Could not parse liquidity from Osmosis-provided position"),
        // note: we would never reach here due to error in get_position if not found, but it is a valid branch in our code
        None => Decimal::zero(),
    };

    let asset0 = position.asset0.expect("Could not find asset0 in position");
    let asset1 = position.asset1.expect("Could not find asset1 in position");
    // should move this into the reply of withdraw position
    let (liquidity_needed_0, liquidity_needed_1) = get_liquidity_needed_for_tokens(
        asset0.amount.clone(),
        asset1.amount.clone(),
        lower_tick,
        upper_tick,
    )?;

    let (deposit, remainders) = get_deposit_amounts_for_liquidity_needed(
        liquidity_needed_0,
        liquidity_needed_1,
        asset0.amount,
        asset1.amount,
    )?;

    let ratio_0_1 = Decimal::from_ratio(deposit.0, deposit.1);
    let (swap_amount, swap_direction) = if !remainders.0.is_zero() {
        (remainders.0, SwapDirection::ZeroToOne)
    } else if !remainders.1.is_zero() {
        (remainders.1, SwapDirection::OneToZero)
    } else {
        // we shouldn't reach here
        (Uint128::zero(), SwapDirection::ZeroToOne)
    };

    // we can naively re-deposit up to however much keeps the proportion of tokens the same. Then swap & re-deposit the proper ratio with the remaining tokens
    let create_position_msg =
        osmosis_std::types::osmosis::concentratedliquidity::v1beta1::MsgCreatePosition {
            pool_id: pool_config.pool_id,
            sender: env.contract.address.to_string(),
            lower_tick: lower_tick
                .try_into()
                .expect("Could not convert lower_tick to i64 from i128"),
            upper_tick: upper_tick
                .try_into()
                .expect("Could not convert upper_tick to i64 from i128"),
            tokens_provided: vec![
                OsmoCoin {
                    denom: pool_config.base_token,
                    amount: deposit.0.to_string(),
                },
                OsmoCoin {
                    denom: pool_config.quote_token,
                    amount: deposit.1.to_string(),
                },
            ],
            // slippage is a mis-nomer here, we won't suffer any slippage. but the pool may still return us more of one of the tokens. This is fine.
            token_min_amount0: with_slippage(deposit.0, vault_config.create_position_max_slippage)?
                .to_string(),
            token_min_amount1: with_slippage(deposit.1, vault_config.create_position_max_slippage)?
                .to_string(),
        };

    // let msg = SubMsg::reply_always(create_position_msg.into(), Replies::CreatePosition.into());

    Ok(Response::default()
        // .add_submessage(msg)
        .add_attribute("action", "execute_rebalance")
        .add_attribute("lower_bound", format!("{:?}", lower_tick))
        .add_attribute("upper_bound", format!("{:?}", upper_tick)))
}

// do create new position
pub fn handle_withdraw_position_response(deps: DepsMut, env: Env, info: MessageInfo) {}

// do swap
pub fn handle_create_position_response(deps: DepsMut, env: Env, info: MessageInfo) {}

// do deposit
pub fn handle_swap_response(deps: DepsMut, env: Env, info: MessageInfo) {}

// do merge position & exit
pub fn handle_deposit_response(deps: DepsMut, env: Env, info: MessageInfo) {}
