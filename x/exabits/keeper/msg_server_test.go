package keeper_test

import (
	"context"
	"testing"

	keepertest "exabits/testutil/keeper"
	"exabits/x/exabits/keeper"
	"exabits/x/exabits/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ExabitsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
