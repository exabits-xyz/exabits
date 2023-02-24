package keeper

import (
	"exabits/x/computing/types"
)

var _ types.QueryServer = Keeper{}
