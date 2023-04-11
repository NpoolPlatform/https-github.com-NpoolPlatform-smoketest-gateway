package testcase

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	modulemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/mgr/module/crud"
	testcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/mgr/testcase/crud"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createModule(ctx context.Context, tx *ent.Tx) error {
	info, err := modulecrud.CreateSet(
		tx.Module.Create(),
		&modulemgrpb.ModuleReq{
			Name: h.ModuleName,
		},
	).Save(ctx)
	if err != nil {
		logger.Sugar().Errorw("createModule", "error", err)
		return err
	}

	moduleID := info.ID.String()
	h.ModuleID = &moduleID
	return nil
}

func (h *createHandler) createTestCase(ctx context.Context, tx *ent.Tx) error {
	info, err := testcasecrud.CreateSet(
		tx.TestCase.Create(),
		&testcasemgrpb.TestCaseReq{
			Name:              h.Name,
			Arguments:         h.Arguments,
			ModuleID:          h.ModuleID,
			ApiID:             h.APIID,
			Description:       h.Description,
			ExpectationResult: h.ExpectationResult,
			TestCaseType:      h.TestCaseType,
		},
	).Save(ctx)
	if err != nil {
		logger.Sugar().Errorw("createTestCase", "error", err)
		return err
	}

	testCaseID := info.ID.String()
	h.ID = &testCaseID
	return nil
}

func (h *Handler) Create(ctx context.Context) (info *npool.TestCase, err error) {
	handler := &createHandler{
		Handler: h,
	}
	// 根据ApiID查询
	// TODO

	// 如果传递了ModuleID则根据ModuleID查询Module
	// TODO

	// 如果未传递ModuleID,传递了ModuleName则创建Module
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createModule(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createTestCase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetTestCase(ctx)
}
