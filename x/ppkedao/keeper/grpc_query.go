package keeper

import (
	"github.com/michaltakac/ppkedao/x/ppkedao/types"
)

var _ types.QueryServer = Keeper{}
