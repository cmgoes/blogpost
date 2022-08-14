package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/ytdbuolf/blogpost/testutil/keeper"
	"github.com/ytdbuolf/blogpost/x/blogpost/keeper"
	"github.com/ytdbuolf/blogpost/x/blogpost/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogpostKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
