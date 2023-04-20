package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) UpdateTestPlan(ctx context.Context) (info *npool.TestPlan, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := testplancrud.UpdateSet(
			cli.TestPlan.UpdateOneID(*h.ID),
			&testplancrud.Req{
				ID:       h.ID,
				Name:     h.Name,
				Executor: h.Executor,
				State:    h.State,
				Deadline: h.Deadline,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	info, err = h.GetTestPlan(ctx)
	if err != nil {
		return nil, err
	}

	return info, nil
}
