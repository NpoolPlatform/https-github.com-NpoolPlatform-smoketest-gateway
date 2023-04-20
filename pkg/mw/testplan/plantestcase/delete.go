package plantestcase

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	plantestcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) DeletePlanTestCase(ctx context.Context) (*npool.PlanTestCase, error) {
	info, err := h.GetPlanTestCase(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := plantestcasecrud.UpdateSet(
			cli.PlanTestCase.UpdateOneID(*h.ID),
			&plantestcasecrud.Req{
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
