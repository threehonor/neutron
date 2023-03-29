package keeper_test

import (
	"github.com/neutron-org/neutron/app"
	"testing"

	testkeeper "github.com/neutron-org/neutron/testutil/cron/keeper"

	"github.com/neutron-org/neutron/x/cron/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	_ = app.GetDefaultConfig()

	k, ctx := testkeeper.CronKeeper(t, nil)
	params := types.Params{
		AdminAddress:    "neutron13xvjxhkkxxhztcugr6weyt76eedj5ucpt4xluv",
		SecurityAddress: "neutron13xvjxhkkxxhztcugr6weyt76eedj5ucpt4xluv",
		Limit:           5,
	}

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
