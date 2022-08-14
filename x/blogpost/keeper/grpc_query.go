package keeper

import (
	"github.com/ytdbuolf/blogpost/x/blogpost/types"
)

var _ types.QueryServer = Keeper{}
