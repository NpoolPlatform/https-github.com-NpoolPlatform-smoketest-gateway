package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
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

func (h *Handler) CreateModule(ctx context.Context) (info *npool.Module, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	id := uuid.New()
	if handler.ID == nil {
		handler.ID = &id
	}

	h.Conds.Name = &cruder.Cond{Op: h.Conds.Name.Op, Val: h.Name}
	if _, err = h.ExistModuleConds(ctx); err != nil {
		return nil, fmt.Errorf("name already exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := modulecrud.CreateSet(
			cli.Module.Create(),
			&modulecrud.Req{
				ID:          h.ID,
				Name:        h.Name,
				Description: h.Description,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetModule(ctx)
}
