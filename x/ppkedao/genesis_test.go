package ppkedao_test

import (
	"testing"

	keepertest "github.com/michaltakac/ppkedao/testutil/keeper"
	"github.com/michaltakac/ppkedao/testutil/nullify"
	"github.com/michaltakac/ppkedao/x/ppkedao"
	"github.com/michaltakac/ppkedao/x/ppkedao/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PpkedaoKeeper(t)
	ppkedao.InitGenesis(ctx, *k, genesisState)
	got := ppkedao.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
