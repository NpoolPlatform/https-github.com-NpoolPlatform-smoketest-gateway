package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	return nil
}

func (h *createHandler) createModule(ctx context.Context, tx *ent.Tx) error {
	info, err := modulecrud.CreateSet(
		tx.Module.Create(),
		&module.ModuleReq{
			Name:        h.Name,
			Description: h.Description,
		},
	).Save(ctx)
	if err != nil {
		return err
	}

	id := info.ID.String()
	h.ID = &id
	return nil
}

func (h *Handler) CreateModule(ctx context.Context) (info *npool.Module, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createModule(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetModule(ctx)
}
