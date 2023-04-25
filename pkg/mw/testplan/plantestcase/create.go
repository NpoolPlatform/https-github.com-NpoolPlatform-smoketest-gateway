package plantestcase

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) CreatePlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	if h.TestCaseResult == nil {
		return nil, fmt.Errorf("invalid testcase result")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
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

	return h.GetPlanTestCase(ctx)
}
