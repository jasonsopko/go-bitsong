package keeper

import (
	"github.com/bitsongofficial/go-bitsong/x/reward"
	rewardTypes "github.com/bitsongofficial/go-bitsong/x/reward/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	cmint "github.com/cosmos/cosmos-sdk/x/mint"
)

// Keeper of the mint store
type Keeper struct {
	bankKeeper   bank.Keeper
	rewardKeeper reward.Keeper
}

// NewKeeper creates a new mint Keeper instance
func NewKeeper(rk reward.Keeper, bk bank.Keeper) Keeper {
	return Keeper{
		rewardKeeper: rk,
		bankKeeper:   bk,
	}
}

func (k Keeper) AddToRewardPool(ctx sdk.Context, coins sdk.Coins) error {
	rewardPool := k.rewardKeeper.GetRewardPool(ctx)
	rewardPool.Amount = rewardPool.Amount.Add(sdk.NewDecCoinsFromCoins(coins...)...)
	k.rewardKeeper.SetRewardPool(ctx, rewardPool)

	return k.bankKeeper.SendCoinsFromModuleToModule(ctx, cmint.ModuleName, rewardTypes.ModuleName, coins)
}

func (k Keeper) GetRewardPoolSupply(ctx sdk.Context) sdk.Coins {
	return k.rewardKeeper.GetRewardPoolSupply(ctx)
}
