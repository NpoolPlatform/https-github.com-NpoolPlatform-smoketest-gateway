package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	plantestcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		_result := h.TestCaseResult.String()
		if _, err := plantestcasecrud.CreateSet(
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
