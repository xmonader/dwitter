package keeper

import (
	"encoding/binary"
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/xmonader/dwitter/x/dwitter/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetTweetsCount(ctx sdk.Context) uint64 {
	byteKey := []byte(types.TweetCountKey)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetTweetsCount(ctx sdk.Context, count uint64) {
	byteKey := []byte(types.TweetCountKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), byteKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) AddTweet(ctx sdk.Context, tweet types.Tweet) (id uint64) {
	count := k.GetTweetsCount(ctx)
	tweet.Id = count
	tweetsKeyByte := []byte(types.TweetKey)
	store := prefix.NewStore(ctx.KVStore(k.storeKey), tweetsKeyByte)

	tweetIdKey := make([]byte, 8)
	binary.BigEndian.PutUint64(tweetIdKey, tweet.Id)

	addedTweetBytes := k.cdc.MustMarshal(&tweet)
	store.Set(tweetIdKey, addedTweetBytes)
	k.SetTweetsCount(ctx, count+1)
	return count

}
