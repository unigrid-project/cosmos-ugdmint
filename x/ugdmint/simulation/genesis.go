package simulation

// DONTCOVER

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/unigrid-project/cosmos-ugdmint/x/ugdmint/types"
)

// Simulation parameter constants
const (
	SubsidyHalvingInterval = "subsidy_halving_interval"
	GoalBonded             = "goal_bonded"
)

var (
	KeySubsidyHalvingInterval = []byte("SubsidyHalvingInterval")
	KeyGoalBonded             = []byte("GoalBonded")
)

// GenSubsidyHalvingInterval randomized subsidy halving interval
func GenSubsidyHalvingInterval(r *rand.Rand) math.LegacyDec {
	return math.LegacyNewDecWithPrec(int64(r.Intn(99)), 0)
}

// GenGoalBonded randomized GoalBonded
func GenGoalBonded(r *rand.Rand) math.LegacyDec {
	return math.LegacyNewDecWithPrec(67, 2)
}

// RandomizedGenState generates a random GenesisState for mint
func RandomizedGenState(simState *module.SimulationState) {
	// minter
	var subsidyHalvingInterval math.LegacyDec
	simState.AppParams.GetOrGenerate(
		string(KeySubsidyHalvingInterval), &subsidyHalvingInterval, simState.Rand,
		func(r *rand.Rand) { subsidyHalvingInterval = GenSubsidyHalvingInterval(r) },
	)

	// params
	var goalBonded math.LegacyDec
	simState.AppParams.GetOrGenerate(
		string(KeyGoalBonded), &goalBonded, simState.Rand,
		func(r *rand.Rand) { goalBonded = GenGoalBonded(r) },
	)

	mintDenom := sdk.DefaultBondDenom
	blocksPerYear := uint64(60 * 60 * 8766 / 5)
	params := types.NewParams(mintDenom, subsidyHalvingInterval, goalBonded, blocksPerYear)

	mintGenesis := types.NewGenesisState(types.InitialMinter(subsidyHalvingInterval), params)

	bz, err := json.MarshalIndent(&mintGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated minting parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(mintGenesis)
}
