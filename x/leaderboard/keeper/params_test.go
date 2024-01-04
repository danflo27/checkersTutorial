package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/leaderboard/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LeaderboardKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
