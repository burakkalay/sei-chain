package migrations

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/sei-protocol/sei-chain/x/dex/types"
)

func PriceSnapshotUpdate(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec, paramStore paramtypes.Subspace) error {
	return nil
}

func migratePriceSnapshotParam(ctx sdk.Context, paramStore paramtypes.Subspace) error {
	defaultParams := types.Params{
		PriceSnapshotRetention: types.DefaultPriceSnapshotRetention,
	}
	paramStore.SetParamSet(ctx, &defaultParams)
	return nil
}