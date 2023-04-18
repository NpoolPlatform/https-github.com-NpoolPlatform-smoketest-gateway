package plantestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	plantestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *createHandler) createPlanTestCas(ctx context.Context, tx *ent.Tx) error {
	info, err := plantestcase.CreateSet(
		tx.PlanTestCase.Create(),
		&plantestcase.PlanTestCasReq{
			TestPlanID:     h.TestCaseID,
			TestCaseID:     h.TestCaseID,
			TestCaseOutput: h.TestCaseOutput,
			TestCaseResult: (*plantestcase.TestCaseResult)(h.TestCaseResult),
			TestUserID:     h.TestUserID,
			RunDuration:    h.RunDuration,
			Description:    h.Description,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	id := info.ID.String()
	h.ID = &id
	return nil
}

func (h *Handler) CreatePlanTestCas(ctx context.Context) (info *npool.PlanTestCas, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createPlanTestCas(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetPlanTestCas(ctx)
}
