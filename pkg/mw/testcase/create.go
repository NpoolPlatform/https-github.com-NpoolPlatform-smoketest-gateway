package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createTestCase(ctx context.Context, cli *ent.Client) error {
	info, err := crud.CreateSet(
		cli.TestCase.Create(),
		&crud.Req{
			EntID:         h.EntID,
			Name:          h.Name,
			Description:   h.Description,
			Input:         h.Input,
			InputDesc:     h.InputDesc,
			ApiID:         h.ApiID,
			ModuleID:      h.ModuleID,
			Expectation:   h.Expectation,
			OutputDesc:    h.OutputDesc,
			TestCaseType:  h.TestCaseType,
			TestCaseClass: h.TestCaseClass,
		},
	).Save(ctx)
	if err != nil {
		return err
	}
	h.ID = &info.ID
	h.EntID = &info.EntID
	return nil
}

func (h *Handler) CreateTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	handler := &createHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.createTestCase(_ctx, cli); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetTestCase(ctx)
}
