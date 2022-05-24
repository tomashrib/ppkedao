package keeper_test

import (
	"testing"

	testkeeper "github.com/michaltakac/ppkedao/testutil/keeper"
	"github.com/michaltakac/ppkedao/x/ppkedao/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PpkedaoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
