package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/xmonader/dwitter/x/dwitter/types"
)

func (k msgServer) CreateTweet(goCtx context.Context, msg *types.MsgCreateTweet) (*types.MsgCreateTweetResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	var tweet = types.Tweet{
		Creator: msg.Creator,
		Content: msg.Content,
	}
	id := k.AddTweet(ctx, tweet)

	return &types.MsgCreateTweetResponse{Id: id}, nil
}
