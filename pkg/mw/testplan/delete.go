package testplan

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) DeleteTestPlan(ctx context.Context) (*npool.TestPlan, error) {
	info, err := h.GetTestPlan(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := testplancrud.UpdateSet(
			cli.TestPlan.UpdateOneID(*h.ID),
			&testplancrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
