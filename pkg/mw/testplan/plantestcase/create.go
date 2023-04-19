package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	plantestcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_result := h.TestCaseResult.String()

		_info, err := plantestcasecrud.CreateSet(
			cli.PlanTestCase.Create(),
			&plantestcasecrud.Req{
				TestPlanID:     h.TestPlanID,
				TestCaseID:     h.TestCaseID,
				TestUserID:     h.TestUserID,
				TestCaseOutput: h.TestCaseOutput,
				Description:    h.Description,
				RunDuration:    h.RunDuration,
				Result:         &_result,
				Index:          &h.Index,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}
		h.ID = &_info.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return h.GetPlanTestCase(ctx)
}
