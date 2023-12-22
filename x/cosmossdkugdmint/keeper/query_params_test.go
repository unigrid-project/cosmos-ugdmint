package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/unigrid-project/cosmos-sdk-ugdmint/testutil/keeper"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

func TestParamsQuery(t *testing.T) {
	keeper, ctx := keepertest.CosmossdkugdmintKeeper(t)
	params := types.DefaultParams()
	require.NoError(t, keeper.SetParams(ctx, params))

	response, err := keeper.Params(ctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
