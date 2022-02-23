package keeper

import (
	"github.com/xmonader/dwitter/x/dwitter/types"
)

var _ types.QueryServer = Keeper{}
