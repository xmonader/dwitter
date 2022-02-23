package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/xmonader/dwitter/x/dwitter/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Tweets(goCtx context.Context, req *types.QueryTweetsRequest) (*types.QueryTweetsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	var tweets []*types.Tweet
	store := ctx.KVStore(k.storeKey)
	tweetsBytesKey := []byte(types.TweetKey)
	tweetsStore := prefix.NewStore(store, tweetsBytesKey)

	pageRes, err := query.Paginate(tweetsStore, req.Pagination, func(key []byte, value []byte) error {
		var tweet types.Tweet
		if err := k.cdc.Unmarshal(value, &tweet); err != nil {
			return err
		}
		tweets = append(tweets, &tweet)
		return nil
	})
	// Throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryTweetsResponse{Tweet: tweets, Pagination: pageRes}, nil
}
