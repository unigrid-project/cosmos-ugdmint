package cosmossdkugdmint_test

import (
	"testing"

	"github.com/unigrid-project/cosmos-sdk-ugdmint/testutil/nullify"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	// k, ctx := keepertest.CosmossdkugdmintKeeper(t)
	// // cosmossdkugdmint.InitGenesis(ctx, k.authKeeper, k, genesisState)
	// got := cosmossdkugdmint.ExportGenesis(ctx, k)
	// require.NotNil(t, got)

	nullify.Fill(&genesisState)
	// nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
