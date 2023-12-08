package module

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
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
		entmodule.FieldEntID,
		entmodule.FieldName,
		entmodule.FieldDescription,
		entmodule.FieldCreatedAt,
		entmodule.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryModule(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Module.Query().Where(entmodule.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entmodule.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entmodule.EntID(*h.EntID))
	}
	h.selectModule(stm)
	return nil
}

func (h *queryHandler) queryModulesByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := modulecrud.SetQueryConds(cli.Module.Query(), h.Conds)
	if err != nil {
		return err
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
		if err := handler.queryModulesByConds(_ctx, cli); err != nil {
			return err
		}

		handler.
			stm.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
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
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}

	return handler.infos[0], nil
}
