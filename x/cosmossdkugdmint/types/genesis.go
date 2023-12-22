package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # genesis/types/import
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// InflationCalculationFn defines the function required to calculate inflation rate during
// BeginBlock. It receives the minter and params stored in the keeper, along with the current
// bondedRatio and returns the newly calculated inflation rate.
// It can be used to specify a custom inflation calculation logic, instead of relying on the
// default logic provided by the sdk.
type InflationCalculationFn func(ctx sdk.Context, minter Minter, params Params, bondedRatio math.LegacyDec) math.LegacyDec

// DefaultInflationCalculationFn is the default function used to calculate inflation.
func DefaultInflationCalculationFn(_ sdk.Context, minter Minter, params Params, bondedRatio math.LegacyDec) math.LegacyDec {
	return math.LegacyNewDec(0)
}

// NewGenesisState creates a new GenesisState object
func NewGenesisState(minter Minter, params Params) *GenesisState {
	return &GenesisState{
		Minter: minter,
		Params: params,
	}
}

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Minter: DefaultInitialMinter(),
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return ValidateMinter(gs.Minter)
}
