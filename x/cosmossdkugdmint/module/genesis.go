package cosmossdkugdmint

import (
	"context"

	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/keeper"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx context.Context, ak types.AccountKeeper, k keeper.Keeper, genState types.GenesisState) {
	k.SetMinter(ctx, genState.Minter)
	if err := k.SetParams(ctx, genState.Params); err != nil {
		panic(err)
	}

	ak.GetModuleAccount(ctx, types.ModuleName)

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.Minter = k.GetMinter(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
