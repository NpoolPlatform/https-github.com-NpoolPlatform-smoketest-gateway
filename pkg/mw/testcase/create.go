package testcase

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
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
	//TODO:Need Refactor
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		modules, err := cli.Module.Query().Where(
			entmodule.Name(*h.ModuleName),
			entmodule.DeletedAt(0)).
			All(_ctx)
		if err != nil {
			return err
		}

		if len(modules) == 0 {
			_entModule := cli.Module.Create()
			_entModule.SetName(*h.ModuleName)
			info, err := _entModule.Save(_ctx)
			if err != nil {
				return err
			}
			h.ModuleID = &info.ID
			return nil
		}

		if len(modules) == 1 {
			id := modules[0].ID
			h.ModuleID = &id
			return nil
		}

		if len(modules) > 1 {
			return fmt.Errorf("too many records")
		}

		return nil
	})

	if err != nil {
		return err
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
