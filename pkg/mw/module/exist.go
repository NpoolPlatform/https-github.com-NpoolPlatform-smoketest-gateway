package module

import (
	"context"
	"fmt"

	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
)

func (h *Handler) ExistModuleByName(ctx context.Context) (exist bool, err error) {
	if h.Name == nil {
		return false, fmt.Errorf("invalid name")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Module.
			Query().
			Where(
				entmodule.Name(*h.Name),
				entmodule.DeletedAt(0),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistModule(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			Module.
			Query().
			Where(
				entmodule.ID(*h.ID),
			).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistModuleConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := modulecrud.SetQueryConds(cli.Module.Query(), h.Conds)
		if err != nil {
			return err
		}

		exist, err = info.Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
