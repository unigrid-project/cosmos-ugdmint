package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/unigrid-project/cosmos-sdk-ugdmint/testutil/keeper"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/keeper"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

func setupMsgServer(t testing.TB) (keeper.Keeper, types.MsgServer, context.Context) {
	k, ctx := keepertest.CosmossdkugdmintKeeper(t)
	return k, keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer(t *testing.T) {
	k, ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
	require.NotEmpty(t, k)
}
