package keeper

import (
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

var _ types.QueryServer = Keeper{}
