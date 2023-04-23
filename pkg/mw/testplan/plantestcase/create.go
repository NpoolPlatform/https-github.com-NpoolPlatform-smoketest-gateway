package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.CreateSet(
			cli.PlanTestCase.Create(),
			&crud.Req{
				TestPlanID:     h.TestPlanID,
				TestCaseID:     h.TestCaseID,
				TestUserID:     h.TestUserID,
				TestCaseOutput: h.TestCaseOutput,
				Description:    h.Description,
				RunDuration:    h.RunDuration,
				Result:         h.TestCaseResult,
				Index:          h.Index,
			},
		).Save(_ctx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return h.GetPlanTestCase(ctx)
}
