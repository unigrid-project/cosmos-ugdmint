package cosmossdkugdmint

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"cosmossdk.io/core/appmodule"
	"cosmossdk.io/core/store"
	"cosmossdk.io/depinject"
	"cosmossdk.io/log"
	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	// this line is used by starport scaffolding # 1

	modulev1 "github.com/unigrid-project/cosmos-sdk-ugdmint/api/cosmossdkugdmint/cosmossdkugdmint/module"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/keeper"
	"github.com/unigrid-project/cosmos-sdk-ugdmint/x/cosmossdkugdmint/types"
)

var (
	_ module.AppModuleBasic      = (*AppModule)(nil)
	_ module.AppModuleSimulation = (*AppModule)(nil)
	_ module.HasGenesis          = (*AppModule)(nil)
	_ module.HasInvariants       = (*AppModule)(nil)
	_ module.HasConsensusVersion = (*AppModule)(nil)

	_ appmodule.AppModule       = (*AppModule)(nil)
	_ appmodule.HasBeginBlocker = (*AppModule)(nil)
	_ appmodule.HasEndBlocker   = (*AppModule)(nil)
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the
// independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string.
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used
// to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore.
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message.
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage.
// The default GenesisState need to be defined by the module developer and is primarily used for testing.
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form.
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module.
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	if err := types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx)); err != nil {
		panic(err)
	}
}

// GetTxCmd returns the root Tx command for the module.
// These commands enrich the AutoCLI tx commands.
// When creating non AutoCLI commands, add the following:
// func (a AppModuleBasic) GetTxCmd() *cobra.Command {
//    return cli.GetTxCmd()
// }

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
	}
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.accountKeeper, am.keeper, genState)
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module.
// It should be incremented on each consensus-breaking change introduced by the module.
// To avoid wrong/empty versions, the initial version should be set to 1.
func (AppModule) ConsensusVersion() uint64 { return 1 }

var (
	prevBlockTime = time.Now()
	account       authtypes.BaseAccount
)

type StatusResponse struct {
	Result struct {
		SyncInfo struct {
			CatchingUp bool `json:"catching_up"`
		} `json:"sync_info"`
	} `json:"result"`
}

// Minting module event types
const (
	AttributeKeyBondedRatio            = "bonded_ratio"
	AttributeKeySubsidyHalvingInterval = "subsidy_halving_interval"
)

// BeginBlock contains the logic that is automatically triggered at the beginning of each block.
// The begin block implementation is optional.
func (am AppModule) BeginBlock(ctx context.Context) error {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	// fetch stored minter & params
	minter := am.keeper.GetMinter(ctx)
	params := am.keeper.GetParams(ctx)
	height := uint64(sdkCtx.BlockHeight())
	bondedRatio, _ := am.keeper.BondedRatio(sdkCtx)

	minter.SubsidyHalvingInterval = params.SubsidyHalvingInterval
	am.keeper.SetMinter(ctx, minter)

	// get the previous block time from the context
	prevCtx := sdk.NewContext(sdkCtx.MultiStore(), sdkCtx.BlockHeader(), false, log.NewNopLogger()).WithBlockHeight(sdkCtx.BlockHeight() - 1)
	// mint coins, update supply
	mintedCoins := minter.BlockProvision(params, height, sdkCtx, prevCtx)
	ok, mintedCoin := mintedCoins.Find("ugd")

	if !ok {
		_, mintedCoin = mintedCoins.Find("fermi")
	}
	err := am.keeper.MintCoins(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	// send the minted coins to the fee collector account
	err = am.keeper.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		panic(err)
	}

	if mintedCoin.Amount.IsInt64() {
		defer telemetry.ModuleSetGauge(types.ModuleName, float32(mintedCoin.Amount.Int64()), "minted_tokens")
	}

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.ModuleName,
			sdk.NewAttribute(AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(AttributeKeySubsidyHalvingInterval, minter.SubsidyHalvingInterval.String()),
			//sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoin.Amount.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoins.String()),
		),
	)

	//Start the mint cache and minting of new tokens when thier are any in hedgehog.
	mc := types.GetCache()
	//fmt.Printf("height: %d\n", height)
	m, mErr := mc.Read(height)
	if mErr == nil {
		//fmt.Println("There were no errors when checking height. its time to mint to address!!")
		acc, aErr := types.ConvertStringToAcc(m.Address)

		if aErr != nil {
			fmt.Println("convert to account failed")
			panic("error!!!!")
		}
		// get the actual account from the account keeper
		account := am.keeper.GetAccount(ctx, acc)
		//fmt.Println("Acc:", acc)

		if account == nil {
			// Create a new BaseAccount with the address
			baseAcc := authtypes.NewBaseAccountWithAddress(acc)
			//fmt.Println("BaseAccount:", baseAcc)
			// Set the initial balance for the account (if you have any initial balance to set)
			// baseAcc.SetCoins(initialBalance)
			//fmt.Println("baseAcc.PubKey:", baseAcc.PubKey)
			// Convert the BaseAccount to a DelayedVestingAccount
			endTime := sdkCtx.BlockTime().Add(10 * 365 * 24 * time.Hour) // 10 years from now
			vestingAcc, _ := vestingtypes.NewDelayedVestingAccount(baseAcc, sdk.Coins{}, endTime.Unix())
			//fmt.Println("Vesting Account:", vestingAcc)
			// Set this new account in the keeper
			am.keeper.SetAccount(ctx, vestingAcc)
		} else if baseAcc, ok := account.(*authtypes.BaseAccount); ok {
			endTime := sdkCtx.BlockTime().Add(10 * 365 * 24 * time.Hour) // 10 years from now
			currentBalances := am.keeper.GetAllBalances(ctx, baseAcc.GetAddress())
			//fmt.Println("baseAcc.PubKey:", baseAcc.PubKey)
			vestingAcc, _ := vestingtypes.NewDelayedVestingAccount(baseAcc, currentBalances, endTime.Unix())
			am.keeper.SetAccount(ctx, vestingAcc)
		} else if baseAcc, ok := account.(*vestingtypes.DelayedVestingAccount); ok {
			currentBalances := am.keeper.GetAllBalances(ctx, baseAcc.GetAddress())

			startTime := sdkCtx.BlockTime().Unix() // Current block time as start time

			// Calculate the amount for each vesting period for each coin in currentBalances
			amountPerPeriod := sdk.Coins{}
			for _, coin := range currentBalances {
				amount := coin.Amount.Quo(math.NewInt(10))
				amountPerPeriod = append(amountPerPeriod, sdk.NewCoin(coin.Denom, amount))
			}

			// Create 10 vesting periods, each 1 minute apart
			periods := vestingtypes.Periods{}
			for i := 0; i < 10; i++ {
				period := vestingtypes.Period{
					Length: 60, // 60 seconds = 1 minute
					Amount: amountPerPeriod,
				}
				periods = append(periods, period)
			}
			fmt.Println("baseAcc.PubKey:", baseAcc.PubKey)
			baseAccount := &authtypes.BaseAccount{
				Address:       baseAcc.Address,
				PubKey:        baseAcc.PubKey,
				AccountNumber: baseAcc.AccountNumber,
				Sequence:      baseAcc.Sequence,
			}

			// Create the PeriodicVestingAccount
			vestingAcc, _ := vestingtypes.NewPeriodicVestingAccount(baseAccount, currentBalances, startTime, periods)
			am.keeper.SetAccount(ctx, vestingAcc)
		} //else if baseAcc, ok := account.(*vestingtypes.PeriodicVestingAccount); ok {
		//}

		coins := types.ConvertIntToCoin(params, m.Amount)
		//fmt.Println("time to mint")
		am.keeper.MintCoins(ctx, coins)
		//fmt.Printf("Coins are minted to address = %s\n", acc.String())
		mErr := am.keeper.AddNewMint(ctx, coins, acc)
		if mErr != nil {
			fmt.Println(mErr.Error())
		}
		//fmt.Println("Coins have been minted")
	}

	return nil
}

// EndBlock contains the logic that is automatically triggered at the end of each block.
// The end block implementation is optional.
func (am AppModule) EndBlock(_ context.Context) error {
	return nil
}

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (am AppModule) IsOnePerModuleType() {}

// IsAppModule implements the appmodule.AppModule interface.
func (am AppModule) IsAppModule() {}

// ----------------------------------------------------------------------------
// App Wiring Setup
// ----------------------------------------------------------------------------

func init() {
	appmodule.Register(
		&modulev1.Module{},
		appmodule.Provide(ProvideModule),
	)
}

type ModuleInputs struct {
	depinject.In

	StoreService           store.KVStoreService
	Cdc                    codec.Codec
	Config                 *modulev1.Module
	Logger                 log.Logger
	InflationCalculationFn types.InflationCalculationFn `optional:"true"`

	AccountKeeper types.AccountKeeper
	BankKeeper    types.BankKeeper
	StakingKeeper *stakingkeeper.Keeper
}

type ModuleOutputs struct {
	depinject.Out

	CosmossdkugdmintKeeper keeper.Keeper
	Module                 appmodule.AppModule
}

func ProvideModule(in ModuleInputs) ModuleOutputs {
	feeCollectorName := in.Config.FeeCollectorName
	if feeCollectorName == "" {
		feeCollectorName = authtypes.FeeCollectorName
	}

	// default to governance authority if not provided
	authority := authtypes.NewModuleAddress(govtypes.ModuleName)
	if in.Config.Authority != "" {
		authority = authtypes.NewModuleAddressOrBech32Address(in.Config.Authority)
	}
	k := keeper.NewKeeper(
		in.Cdc,
		in.StoreService,
		in.Logger,
		*in.StakingKeeper,
		in.AccountKeeper,
		in.BankKeeper,
		feeCollectorName,
		authority.String(),
	)
	m := NewAppModule(
		in.Cdc,
		k,
		in.AccountKeeper,
		in.BankKeeper,
	)

	return ModuleOutputs{CosmossdkugdmintKeeper: k, Module: m}
}
