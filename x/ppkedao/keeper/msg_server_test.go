package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/michaltakac/ppkedao/testutil/keeper"
	"github.com/michaltakac/ppkedao/x/ppkedao/keeper"
	"github.com/michaltakac/ppkedao/x/ppkedao/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PpkedaoKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
