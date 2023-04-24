package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	modulecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
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
	if exist, _ := modulecli.ExistModuleByName(ctx, *h.Name); exist {
		// GetModuleByName
	}

	conds := &modulecrud.Conds{}
	conds.Name = &cruder.Cond{Op: conds.Name.Op, Val: h.ModuleName}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := modulecrud.SetQueryConds(cli.Module.Query(), conds)
		if err != nil {
			return err
		}

		_, err = info.Exist(_ctx)
		return err
	})
	if err != nil {
		return err
	}

	_module, err := modulecrud.CreateSet(&ent.ModuleCreate{}, &modulecrud.Req{
		Name: h.ModuleName,
	}).Save(ctx)
	if err != nil {
		return err
	}

	h.ModuleID = &_module.ID
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
