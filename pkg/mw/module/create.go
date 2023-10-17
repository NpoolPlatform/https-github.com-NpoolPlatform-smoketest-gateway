package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) CreateModule(ctx context.Context) (info *npool.Module, err error) {
	if exist, err := h.ExistModuleByName(ctx); err != nil || exist {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.Module.Create(),
			&crud.Req{
				ID:          h.ID,
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
