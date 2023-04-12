package module

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/module"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entModule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.ModuleSelect
	infos []*npool.Module
	total uint32
}

func (h *queryHandler) selectModule(stm *ent.ModuleQuery) {
	h.stm = stm.Select(
		entModule.FieldID,
		entModule.FieldName,
		entModule.FieldDescription,
		entModule.FieldCreatedAt,
		entModule.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) queryModule(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid module id")
	}
	h.selectModule(
		cli.Module.
			Query().
			Where(
				entModule.ID(uuid.MustParse(*h.ID)),
				entModule.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryModuleByConds(ctx context.Context, cli *ent.Client) (err error) {
	if h.Conds == nil {
		return fmt.Errorf("invalid conds")
	}

	stm := cli.Module.Query()
	if h.Conds.ID != nil {
		stm = stm.Where(
			entModule.ID(uuid.MustParse(h.Conds.GetID().GetValue())),
		)
	}
	if h.Conds.Name != nil {
		stm = stm.Where(
			entModule.Name(h.Conds.GetName().GetValue()),
		)
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
		if err := handler.queryModuleByConds(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
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
