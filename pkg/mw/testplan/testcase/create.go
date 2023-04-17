package planrelatedtestcase

import (
	"context"

	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/testcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	planrelatedtestcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/mgr/testplan/testcase/crud"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *createHandler) createPlanRelatedTestCase(ctx context.Context, tx *ent.Tx) error {
	info, err := planrelatedtestcasecrud.CreateSet(
		tx.PlanRelatedTestCase.Create(),
		&planrelatedtestcase.PlanRelatedTestCaseReq{
			TestPlanID:     h.TestCaseID,
			TestCaseID:     h.TestCaseID,
			TestCaseOutput: h.TestCaseOutput,
			TestCaseResult: (*planrelatedtestcase.TestCaseResult)(h.TestCaseResult),
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

func (h *Handler) CreatePlanRelatedTestCase(ctx context.Context) (info *npool.PlanRelatedTestCase, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createPlanRelatedTestCase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetPlanRelatedTestCase(ctx)
}
