package plantestcase

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	testcasemwpb "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	entmodule "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	entplantestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
	enttestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.PlanTestCaseSelect
	stmCount  *ent.PlanTestCaseSelect
	infos     []*npool.PlanTestCase
	total     uint32
}

func (h *queryHandler) selectPlanTestCase(stm *ent.PlanTestCaseQuery) *ent.PlanTestCaseSelect {
	return stm.Select(entplantestcase.FieldID)
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entplantestcase.Table)
	s.LeftJoin(t).
		On(
			s.C(entplantestcase.FieldID),
			t.C(entplantestcase.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entplantestcase.FieldEntID), "ent_id"),
			sql.As(t.C(entplantestcase.FieldTestPlanID), "test_plan_id"),
			sql.As(t.C(entplantestcase.FieldTestCaseID), "test_case_id"),
			sql.As(t.C(entplantestcase.FieldInput), "input"),
			sql.As(t.C(entplantestcase.FieldOutput), "output"),
			sql.As(t.C(entplantestcase.FieldResult), "result"),
			sql.As(t.C(entplantestcase.FieldTestUserID), "test_user_id"),
			sql.As(t.C(entplantestcase.FieldIndex), "index"),
			sql.As(t.C(entplantestcase.FieldRunDuration), "run_duration"),
			sql.As(t.C(entplantestcase.FieldDescription), "description"),
			sql.As(t.C(entplantestcase.FieldCreatedAt), "created_at"),
			sql.As(t.C(entplantestcase.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinTestCase(s *sql.Selector) {
	t1 := sql.Table(enttestcase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entplantestcase.FieldTestCaseID),
			t1.C(enttestcase.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(enttestcase.FieldName), "test_case_name"),
			sql.As(t1.C(enttestcase.FieldAPIID), "test_case_api_id"),
			sql.As(t1.C(enttestcase.FieldTestCaseType), "test_case_type"),
			sql.As(t1.C(enttestcase.FieldTestCaseClass), "test_case_class"),
		)
	t2 := sql.Table(entmodule.Table)
	s.LeftJoin(t2).
		On(
			t1.C(enttestcase.FieldModuleID),
			t2.C(entmodule.FieldEntID),
		).
		AppendSelect(
			sql.As(t2.C(entmodule.FieldEntID), "module_id"),
			sql.As(t2.C(entmodule.FieldName), "module_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTestCase(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinTestCase(s)
	})
}

func (h *queryHandler) queryPlanTestCase(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid plantestcase id")
	}
	stm := cli.PlanTestCase.Query().Where(entplantestcase.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entplantestcase.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entplantestcase.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPlanTestCase(stm)
	return nil
}

func (h *queryHandler) queryPlanTestCases(cli *ent.Client) (*ent.PlanTestCaseSelect, error) {
	stm, err := crud.SetQueryConds(cli.PlanTestCase.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectPlanTestCase(stm), nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.Result = npool.TestCaseResult(npool.TestCaseResult_value[info.ResultStr])
		info.TestCaseType = testcasemwpb.TestCaseType(testcasemwpb.TestCaseType_value[info.TestCaseTypeStr])
		info.TestCaseClass = testcasemwpb.TestCaseClass(testcasemwpb.TestCaseClass_value[info.TestCaseClassStr])
	}
}

func (h *Handler) GetPlanTestCases(ctx context.Context) ([]*npool.PlanTestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryPlanTestCases(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryPlanTestCases(cli)
		if err != nil {
			return err
		}
		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
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

func (h *Handler) GetPlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPlanTestCase(cli); err != nil {
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
		return nil, fmt.Errorf("plantestcase not exist")
	}

	handler.formalize()
	return handler.infos[0], nil
}
