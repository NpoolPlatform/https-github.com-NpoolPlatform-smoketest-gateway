package testcase

import (
	"context"
	"fmt"

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
	if h.ModuleID == nil {
		return fmt.Errorf("invalid module id")
	}

	if _, err := modulecrud.CreateSet(
		tx.Module.Create(),
		&modulemgrpb.ModuleReq{
			Name: h.ModuleName,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *createHandler) createTestCase(ctx context.Context, tx *ent.Tx) error {
	if _, err := testcasecrud.CreateSet(
		tx.TestCase.Create(),
		&testcasemgrpb.TestCaseReq{
			// ModuleID TODO
			Name:              h.Name,
			Arguments:         h.Arguments,
			ApiID:             h.ApiID,
			ExpectationResult: h.ExpectationResult,
			TestCaseType:      h.TestCaseType,
		},
	).Save(ctx); err != nil {
		return err
	}

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


