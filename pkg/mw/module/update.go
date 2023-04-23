package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) UpdateModule(ctx context.Context) (info *npool.Module, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Module.UpdateOneID(*h.ID),
			&crud.Req{
				ID:          h.ID,
				Name:        h.Name,
				Description: h.Description,
			},
		).Save(_ctx); err != nil {
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
