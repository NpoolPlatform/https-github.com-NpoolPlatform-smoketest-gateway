package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/module"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
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

func (h *Handler) CreateModule(ctx context.Context) (info *npool.Module, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	if _, err := modulecli.ExistModuleByName(ctx, *handler.Name); err == nil {
		return nil, fmt.Errorf("name already exist")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.Module.Create(),
			&crud.Req{
				Name:        h.Name,
				Description: h.Description,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetModule(ctx)
}
