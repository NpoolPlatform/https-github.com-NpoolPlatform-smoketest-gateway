package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool1 "github.com/NpoolPlatform/message/npool"
	modulemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	modulecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	const leastNameLen = 2
	if len(*h.Name) < leastNameLen {
		return fmt.Errorf("name %v too short", *h.Name)
	}
	return nil
}

func (h *createHandler) createModule(ctx context.Context) error {
	modules, _, err := modulecli.GetModuleConds(ctx, &modulemgrpb.Conds{
		Name: &npool1.StringVal{Op: cruder.EQ, Value: *h.ModuleName},
	})
	if err != nil {
		return err
	}

	if len(modules) == 0 {
		module, err := modulecli.CreateModule(ctx, &modulemgrpb.ModuleReq{
			Name: h.ModuleName,
		})
		if err != nil {
			return err
		}
		id := module.ID
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.ModuleID = &_id
	}

	if len(modules) == 1 {
		id := modules[0].ID
		_id, err := uuid.Parse(id)
		if err != nil {
			return err
		}
		h.ModuleID = &_id
	}

	if len(modules) > 1 {
		logger.Sugar().Info(
			"CreateModule",
			"Req", modules,
			"Info", "too many records",
		)
	}

	return nil
}

func (h *createHandler) createTestCase(ctx context.Context) error {
	err := db.WithClient(ctx, func(ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.TestCase.Create(),
			&crud.Req{
				Name:         h.Name,
				Description:  h.Description,
				Input:        h.Input,
				InputDesc:    h.InputDesc,
				ApiID:        h.ApiID,
				ModuleID:     h.ModuleID,
				Expectation:  h.Expectation,
				TestCaseType: h.TestCaseType,
			},
		).Save(ctx)
		if err != nil {
			return err
		}
		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	handler := &createHandler{
		Handler: h,
	}
	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createModule(_ctx); err != nil {
			return err
		}
		if err := handler.createTestCase(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetTestCase(ctx)
}
