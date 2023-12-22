package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/unigrid-project/cosmos-sdk-ugdmint/testutil/keeper"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.CosmossdkugdmintKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
