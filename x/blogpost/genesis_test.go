package blogpost_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/ytdbuolf/blogpost/testutil/keeper"
	"github.com/ytdbuolf/blogpost/testutil/nullify"
	"github.com/ytdbuolf/blogpost/x/blogpost"
	"github.com/ytdbuolf/blogpost/x/blogpost/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogpostKeeper(t)
	blogpost.InitGenesis(ctx, *k, genesisState)
	got := blogpost.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
