package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) CreateTestPlan(ctx context.Context) (info *npool.TestPlan, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.TestPlan.Create(),
			&crud.Req{
				ID:        h.ID,
				Name:      h.Name,
				CreatedBy: h.CreatedBy,
				Executor:  h.Executor,
				Deadline:  h.Deadline,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return h.GetTestPlan(ctx)
}
