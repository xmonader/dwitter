package dwitter_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/xmonader/dwitter/testutil/keeper"
	"github.com/xmonader/dwitter/testutil/nullify"
	"github.com/xmonader/dwitter/x/dwitter"
	"github.com/xmonader/dwitter/x/dwitter/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DwitterKeeper(t)
	dwitter.InitGenesis(ctx, *k, genesisState)
	got := dwitter.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
