package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/ytdbuolf/blogpost/testutil/keeper"
	"github.com/ytdbuolf/blogpost/x/blogpost/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlogpostKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
