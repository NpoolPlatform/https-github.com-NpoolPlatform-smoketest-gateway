package testplan

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/mgr/testplan/crud"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	if h.OwnerID == nil {
		return fmt.Errorf("invalid owner id")
	}
	return nil
}

func (h *createHandler) createTestPlan(ctx context.Context, tx *ent.Tx) error {
	info, err := testplancrud.CreateSet(
		tx.TestPlan.Create(),
		&testplan.TestPlanReq{
			Name:              h.Name,
			OwnerID:           h.OwnerID,
			ResponsibleUserID: h.ResponsibleUserID,
			Deadline:          h.Deadline,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	id := info.ID.String()
	h.ID = &id
	return nil
}

func (h *Handler) CreateTestPlan(ctx context.Context) (info *npool.TestPlan, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createTestPlan(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTestPlan(ctx)
}
