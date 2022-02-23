package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/xmonader/dwitter/testutil/keeper"
	"github.com/xmonader/dwitter/x/dwitter/keeper"
	"github.com/xmonader/dwitter/x/dwitter/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DwitterKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
