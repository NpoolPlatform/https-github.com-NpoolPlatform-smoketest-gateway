package cond

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entcond "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/cond"
)

type queryHandler struct {
	*Handler
	stm   *ent.CondSelect
	infos []*npool.Cond
	total uint32
}

func (h *queryHandler) selectCond(stm *ent.CondQuery) {
	h.stm = stm.Select(
		entcond.FieldID,
		entcond.FieldTestCaseID,
		entcond.FieldCondTestCaseID,
		entcond.FieldCondType,
		entcond.FieldIndex,
		entcond.FieldArgumentMap,
		entcond.FieldCreatedAt,
		entcond.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryCond(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid module id")
	}
	h.selectCond(
		cli.Cond.
			Query().
			Where(
				entcond.ID(*h.ID),
				entcond.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryCondsByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.Cond.Query(), h.Conds)
	if err != nil {
		return err
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectCond(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.CondType = npool.CondType(npool.CondType_value[info.CondTypeStr])
	}
}

func (h *Handler) GetConds(ctx context.Context) ([]*npool.Cond, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCondsByConds(_ctx, cli); err != nil {
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
	handler.formalize()
	return handler.infos, handler.total, nil
}

func (h *Handler) GetCond(ctx context.Context) (info *npool.Cond, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryCond(cli); err != nil {
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
	handler.formalize()
	return handler.infos[0], nil
}
