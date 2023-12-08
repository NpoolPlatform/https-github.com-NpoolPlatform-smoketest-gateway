package testcase

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	enttestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
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
		enttestcase.FieldEntID,
		enttestcase.FieldName,
		enttestcase.FieldDescription,
		enttestcase.FieldAPIID,
		enttestcase.FieldModuleID,
		enttestcase.FieldTestCaseType,
		enttestcase.FieldTestCaseClass,
		enttestcase.FieldInput,
		enttestcase.FieldInputDesc,
		enttestcase.FieldExpectation,
		enttestcase.FieldOutputDesc,
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
			t.C(entmodule.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entmodule.FieldEntID), "module_id"),
			sql.As(t.C(entmodule.FieldName), "module_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {
		h.queryJoinModule(s)
	})
}

func (h *queryHandler) queryTestCase(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid testcase id")
	}
	stm := cli.TestCase.Query().Where(enttestcase.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttestcase.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttestcase.EntID(*h.EntID))
	}
	h.selectTestCase(stm)
	return nil
}

func (h *queryHandler) queryTestCaseByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := crud.SetQueryConds(cli.TestCase.Query(), h.Conds)
	if err != nil {
		return err
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

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TestCaseType = npool.TestCaseType(npool.TestCaseType_value[info.TestCaseTypeStr])
		info.TestCaseClass = npool.TestCaseClass(npool.TestCaseClass_value[info.TestCaseClassStr])
	}
}

func (h *Handler) GetTestCases(ctx context.Context) ([]*npool.TestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTestCaseByConds(ctx, cli); err != nil {
			return err
		}
		handler.queryJoin()
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

func (h *Handler) GetTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	handler := &queryHandler{
		Handler: h,
	}

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
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("testCase not exist")
	}

	handler.formalize()
	return handler.infos[0], nil
}
