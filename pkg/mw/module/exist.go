package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
)

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
		if h.Conds != nil {
			conds := &modulecrud.Conds{}
			if h.Conds.ID != nil {
				conds.ID = &cruder.Cond{Op: h.Conds.ID.Op, Val: h.Conds.ID.Value}
			}
			if h.Conds.Name != nil {
				conds.Name = &cruder.Cond{Op: h.Conds.Name.Op, Val: h.Conds.Name.Value}
			}
			if h.Conds.IDs != nil {
				conds.IDs = &cruder.Cond{Op: h.Conds.ID.Op, Val: h.Conds.IDs.Value}
			}

			q := &ent.ModuleQuery{}
			info, err := modulecrud.SetQueryConds(q, conds)
			if err != nil {
				return err
			}

			_, err = info.Exist(_ctx)
			return err
		}
		return fmt.Errorf("invalid conds")
	})
	if err != nil {
		return false, err
	}
	return true, nil
}
