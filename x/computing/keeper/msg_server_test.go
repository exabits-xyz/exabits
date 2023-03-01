package keeper_test

import (
	"context"
	"testing"

	keepertest "exabits/testutil/keeper"
	"exabits/x/computing/keeper"
	"exabits/x/computing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ComputingKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
