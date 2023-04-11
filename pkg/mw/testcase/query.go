package testcase

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	enttestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.TestCaseSelect
	infos []*npool.TestCase
	total uint32
}

func (h *queryHandler) selectTestCase(stm *ent.TestCaseQuery) {
	h.stm = stm.Select(
		enttestcase.FieldID,
		enttestcase.FieldName,
		enttestcase.FieldDescription,
		enttestcase.FieldModuleID,
		enttestcase.FieldArguments,
		enttestcase.FieldExpectationResult,
		enttestcase.FieldDeprecated,
		enttestcase.FieldCreatedAt,
		enttestcase.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoinModule(s *sql.Selector) {
	t := sql.Table(entmodule.Table)
	s.LeftJoin(t).
		On(
			s.C(enttestcase.FieldModuleID),
			t.C(entmodule.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entmodule.FieldID), "module_id"),
			sql.As(t.C(entmodule.FieldName), "module_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinModule(s)
	})
}

func (h *queryHandler) queryTestCase(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid testcase id")
	}
	h.selectTestCase(
		cli.TestCase.
			Query().
			Where(
				enttestcase.ID(uuid.MustParse(*h.ID)),
				enttestcase.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryTestCaseByConds(ctx context.Context, cli *ent.Client) (err error) {
	if h.Conds == nil {
		return fmt.Errorf("invalid conds")
	}

	stm := cli.TestCase.Query()
	if h.Conds.ID != nil {
		stm = stm.Where(
			enttestcase.ID(uuid.MustParse(h.Conds.GetID().GetValue())),
		)
	}

	if h.Conds.Deprecated != nil {
		stm = stm.Where(
			enttestcase.Deprecated(h.Conds.GetDeprecated().GetValue()),
		)
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectTestCase(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTestCases(ctx context.Context, cli *ent.Client) ([]*npool.TestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTestCaseByConds(ctx, cli); err != nil {
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

func (h *Handler) GetTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	fmt.Println("ID: ", *handler.ID)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTestCase(cli); err != nil {
			return err
		}
		handler.queryJoin()
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
