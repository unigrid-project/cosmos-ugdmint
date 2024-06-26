package ugdmint

import (
	"context"

	"fmt"
	"time"

	"cosmossdk.io/log"
	"cosmossdk.io/math"

	"runtime/debug"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	"github.com/unigrid-project/cosmos-ugdmint/x/ugdmint/keeper"
	"github.com/unigrid-project/cosmos-ugdmint/x/ugdmint/types"
)

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

// BeginBlocker mints new tokens for the previous block.
func BeginBlocker(goCtx context.Context, k keeper.Keeper) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recovered from panic in BeginBlocker: %v\n", r)
			debug.PrintStack() // Print the stack trace
		}
	}()

	ctx := sdk.UnwrapSDKContext(goCtx)
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyBeginBlocker)
	fmt.Println("BeginBlocker: Starting")

	// Fetch stored minter & params
	minter := k.GetMinter(ctx)
	if (minter == types.Minter{}) {
		fmt.Println("BeginBlocker: Minter is empty")
		return
	}

	params := k.GetParams(ctx)
	if (params == types.Params{}) {
		fmt.Println("BeginBlocker: Params are empty")
		return
	}

	height := uint64(ctx.BlockHeight())
	fmt.Printf("BeginBlocker: Current block height: %d\n", height)

	bondedRatio, err := k.BondedRatio(ctx)
	if err != nil {
		fmt.Println("BeginBlocker: Error getting bonded ratio:", err)
		return
	}

	minter.SubsidyHalvingInterval = params.SubsidyHalvingInterval
	k.SetMinter(goCtx, minter)

	// Get the previous block time from the context
	prevCtx := sdk.NewContext(ctx.MultiStore(), ctx.BlockHeader(), false, log.NewNopLogger()).WithBlockHeight(ctx.BlockHeight() - 1)
	mintedCoins := minter.BlockProvision(params, height, ctx, prevCtx)
	if mintedCoins.Empty() {
		fmt.Println("BeginBlocker: Minted coins are empty")
		return
	}

	err2 := k.MintCoins(goCtx, mintedCoins)
	if err2 != nil {
		fmt.Println("BeginBlocker: Error minting coins:", err2)
		return
	}

	err = k.AddCollectedFees(ctx, mintedCoins)
	if err != nil {
		fmt.Println("BeginBlocker: Error adding collected fees:", err)
		return
	}

	fmt.Println("MintedCoins:", mintedCoins)
	for _, coin := range mintedCoins {
		if coin.Amount.BigInt() != nil && coin.Amount.IsInt64() {
			defer telemetry.ModuleSetGauge(types.ModuleName, float32(coin.Amount.Int64()), "minted_tokens")
		} else {
			fmt.Println("BeginBlocker: coin.Amount is nil or not an int64 for coin", coin.Denom)
		}
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUGDMint,
			sdk.NewAttribute(types.AttributeKeyBondedRatio, bondedRatio.String()),
			sdk.NewAttribute(types.AttributeKeySubsidyHalvingInterval, minter.SubsidyHalvingInterval.String()),
			sdk.NewAttribute(sdk.AttributeKeyAmount, mintedCoins.String()),
		),
	)

	// Initialize and start the mint cache
	fmt.Println("BeginBlocker: Initializing mint cache")
	mc := types.GetCache()
	if mc == nil {
		fmt.Println("BeginBlocker: Mint cache is nil")
		return
	}

	// Use mc.Read to check if mint data exists for the current height
	mint, err := mc.Read(height)
	if err != nil {
		fmt.Printf("BeginBlocker: No mint data for current block height %d: %v\n", height, err)
		fmt.Println("BeginBlocker: No mint data available. Skipping minting process.")
		return
	} else {
		// Process the mint if it exists for the current height
		fmt.Println("BeginBlocker: Mint data found for current block height")
		acc, aErr := types.ConvertStringToAcc(mint.Address)
		if aErr != nil {
			fmt.Println("BeginBlocker: Error converting string to account:", aErr)
			return
		}

		account := k.GetAccount(ctx, acc)
		fmt.Println("BeginBlocker: Retrieved account:", account)

		if account == nil {
			baseAcc := authtypes.NewBaseAccountWithAddress(acc)
			fmt.Println("BeginBlocker: Created new base account:", baseAcc)

			accNum, err := k.GetNextAccountNumber(ctx)
			if err != nil {
				fmt.Println("BeginBlocker: Error getting next account number:", err)
				return
			}

			baseAcc.SetAccountNumber(accNum)
			endTime := ctx.BlockTime().Add(10 * 365 * 24 * time.Hour) // 10 years from now

			// Mint the coins first and include them in the original vesting
			coins := sdk.NewCoins(sdk.NewCoin("uugd", math.NewInt(int64(mint.Amount))))
			if err := k.MintCoins(goCtx, coins); err != nil {
				fmt.Println("BeginBlocker: Error minting coins:", err)
				return
			}

			vestingAcc, err := vestingtypes.NewDelayedVestingAccount(baseAcc, coins, endTime.Unix())
			if err != nil {
				fmt.Println("BeginBlocker: Error creating vesting account:", err)
				return
			}
			fmt.Println("BeginBlocker: Created new vesting account:", vestingAcc)

			if err := k.SetAccount(ctx, vestingAcc); err != nil {
				fmt.Println("BeginBlocker: Error setting account:", err)
				return
			}
		} else if baseAcc, ok := account.(*authtypes.BaseAccount); ok {
			endTime := ctx.BlockTime().Add(10 * 365 * 24 * time.Hour) // 10 years from now
			currentBalances := k.GetAllBalances(ctx, baseAcc.GetAddress())
			if currentBalances == nil {
				fmt.Println("BeginBlocker: Current balances are nil")
				return
			}

			vestingAcc, err := vestingtypes.NewDelayedVestingAccount(baseAcc, currentBalances, endTime.Unix())
			if err != nil {
				fmt.Println("BeginBlocker: Error creating vesting account:", err)
				return
			}
			if err := k.SetAccount(ctx, vestingAcc); err != nil {
				fmt.Println("BeginBlocker: Error setting account:", err)
				return
			}
		} else if baseAcc, ok := account.(*vestingtypes.DelayedVestingAccount); ok {
			currentBalances := k.GetAllBalances(ctx, baseAcc.GetAddress())
			if currentBalances == nil {
				fmt.Println("BeginBlocker: Current balances are nil")
				return
			}

			startTime := ctx.BlockTime().Unix()
			amountPerPeriod := sdk.Coins{}
			for _, coin := range currentBalances {
				amount := coin.Amount.Quo(math.NewInt(10))
				amountPerPeriod = append(amountPerPeriod, sdk.NewCoin(coin.Denom, amount))
			}

			periods := vestingtypes.Periods{}
			for i := 0; i < 10; i++ {
				period := vestingtypes.Period{
					Length: 60,
					Amount: amountPerPeriod,
				}
				periods = append(periods, period)
			}

			baseAccount := &authtypes.BaseAccount{
				Address:       baseAcc.Address,
				PubKey:        baseAcc.PubKey,
				AccountNumber: baseAcc.AccountNumber,
				Sequence:      baseAcc.Sequence,
			}
			vestingAcc, err := vestingtypes.NewPeriodicVestingAccount(baseAccount, currentBalances, startTime, periods)
			if err != nil {
				fmt.Println("BeginBlocker: Error creating periodic vesting account:", err)
				return
			}
			if err := k.SetAccount(ctx, vestingAcc); err != nil {
				fmt.Println("BeginBlocker: Error setting account:", err)
				return
			}
		}

		// Ensure the coins are converted to 'uugd'
		coins := sdk.NewCoins(sdk.NewCoin("uugd", math.NewInt(int64(mint.Amount))))
		if coins.Empty() {
			fmt.Println("BeginBlocker: Coins conversion resulted in empty coins")
			return
		}

		if err := k.MintCoins(goCtx, coins); err != nil {
			fmt.Println("BeginBlocker: Error minting coins:", err)
			return
		}

		if err := k.AddNewMint(ctx, coins, acc); err != nil {
			fmt.Println("BeginBlocker: Error adding new mint:", err)
			return
		}

		mintRecord := types.MintRecord{
			BlockHeight: ctx.BlockHeight(),
			Account:     "target_account_address", // Replace with actual account
			Amount:      mintedCoins,
		}

		if err := k.SetMintRecord(ctx, mintRecord); err != nil {
			fmt.Println("BeginBlocker: Error storing mints in the KVstore:", err)
			return
		}

		fmt.Println("BeginBlocker: Mint process completed successfully")
	}

	fmt.Println("BeginBlocker: Completed successfully")
}
