package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/xmonader/dwitter/testutil/keeper"
	"github.com/xmonader/dwitter/x/dwitter/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.DwitterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
