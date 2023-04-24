package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) UpdateModule(ctx context.Context) (info *npool.Module, err error) {
	info, err = h.GetModule(ctx)
	if err != nil {
		return nil, err
	}

	// determine whether name is the same as before
	req := &crud.Req{}
	if h.Name != nil {
		if info.Name != *h.Name {
			if _, err := h.ExistModuleByName(ctx); err != nil {
				return nil, fmt.Errorf("name already exist")
			}
			req.Name = h.Name
		}
	}
	req.Description = h.Description

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(cli.Module.UpdateOneID(*h.ID), req).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	info, err = h.GetModule(ctx)
	if err != nil {
		return nil, err
	}

	return info, nil
}
