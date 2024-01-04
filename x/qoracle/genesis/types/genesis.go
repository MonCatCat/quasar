package types

import (
	qosmotypes "github.com/MonCatCat/quasar/x/qoracle/osmosis/types"
	types "github.com/MonCatCat/quasar/x/qoracle/types"
	host "github.com/cosmos/ibc-go/v4/modules/core/24-host"
)

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:              types.DefaultParams(),
		OsmosisGenesisState: DefaultOsmosisGenesis(),
	}
}

// NewGenesisState creates and returns a new GenesisState instance from the provided controller and host genesis state types
func NewGenesisState(params types.Params, osmosisGenesisState OsmosisGenesisState) *GenesisState {
	return &GenesisState{
		Params:              params,
		OsmosisGenesisState: osmosisGenesisState,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	if err := gs.OsmosisGenesisState.Validate(); err != nil {
		return err
	}

	return nil
}

// DefaultOsmosisGenesis creates and returns the default qoracle DefaultOsmosisGenesis
func DefaultOsmosisGenesis() OsmosisGenesisState {
	return OsmosisGenesisState{
		Port:   qosmotypes.PortID,
		Params: qosmotypes.DefaultParams(),
	}
}

// NewOsmosisGenesisState creates a returns a new OsmosisGenesisState instance
func NewOsmosisGenesisState(port string, params qosmotypes.Params) OsmosisGenesisState {
	return OsmosisGenesisState{
		Port:   port,
		Params: params,
	}
}

// Validate performs basic validation of the OsmosisGenesisState
func (gs OsmosisGenesisState) Validate() error {

	if err := host.PortIdentifierValidator(gs.Port); err != nil {
		return err
	}

	if err := gs.Params.Validate(); err != nil {
		return err
	}

	return nil
}
