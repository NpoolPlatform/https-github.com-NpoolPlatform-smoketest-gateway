package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"
	modulecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
)

type queryHandler struct {
	*Handler
	stm   *ent.ModuleSelect
	infos []*npool.Module
	total uint32
}

func (h *queryHandler) selectModule(stm *ent.ModuleQuery) {
	h.stm = stm.Select(
		entmodule.FieldID,
		entmodule.FieldName,
		entmodule.FieldDescription,
		entmodule.FieldCreatedAt,
		entmodule.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryModule(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid module id")
	}
	h.selectModule(
		cli.Module.
			Query().
			Where(
				entmodule.ID(*h.ID),
				entmodule.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryModules(ctx context.Context, cli *ent.Client) (err error) {
	stm := cli.Module.Query()
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
		stm, err = modulecrud.SetQueryConds(stm, conds)
		if err != nil {
			return err
		}
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectModule(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetModules(ctx context.Context) ([]*npool.Module, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryModules(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(*h.Offset)).
			Limit(int(*h.Limit))
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}

func (h *Handler) GetModule(ctx context.Context) (info *npool.Module, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryModule(cli); err != nil {
			return err
		}
		if err := handler.scan(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	return handler.infos[0], nil
}
