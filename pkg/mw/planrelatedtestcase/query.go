package planrelatedtestcase

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/planrelatedtestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	planrelatedtestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/planrelatedtestcase"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.PlanRelatedTestCaseSelect
	infos []*npool.PlanRelatedTestCase
	total uint32
}

func (h *queryHandler) selectPlanRelatedTestCase(stm *ent.PlanRelatedTestCaseQuery) {
	h.stm = stm.Select(
		planrelatedtestcase.FieldID,
		planrelatedtestcase.FieldTestPlanID,
		planrelatedtestcase.FieldTestCaseID,
		planrelatedtestcase.FieldTestCaseOutput,
		// planrelatedtestcase.FieldTestCaseResult,
		planrelatedtestcase.FieldTestUserID,
		planrelatedtestcase.FieldRunDuration,
		planrelatedtestcase.FieldDescription,
		planrelatedtestcase.FieldCreatedAt,
		planrelatedtestcase.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) queryPlanRelatedTestCase(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid planrelatedtestcase id")
	}
	h.selectPlanRelatedTestCase(
		cli.PlanRelatedTestCase.
			Query().
			Where(
				planrelatedtestcase.ID(uuid.MustParse(*h.ID)),
				planrelatedtestcase.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryPlanRelatedTestCaseByConds(ctx context.Context, cli *ent.Client) (err error) {
	if h.Conds == nil {
		return fmt.Errorf("invalid conds")
	}

	stm := cli.PlanRelatedTestCase.Query()
	if h.Conds.ID != nil {
		stm = stm.Where(
			planrelatedtestcase.ID(uuid.MustParse(h.Conds.GetID().GetValue())),
		)
	}
	if h.Conds.TestPlanID != nil {
		stm = stm.Where(
			planrelatedtestcase.TestPlanID(uuid.MustParse(h.Conds.GetTestPlanID().GetValue())),
		)
	}
	if h.Conds.TestUserID != nil {
		stm = stm.Where(
			planrelatedtestcase.TestUserID(uuid.MustParse(h.Conds.GetTestPlanID().GetValue())),
		)
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectPlanRelatedTestCase(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetPlanRelatedTestCases(ctx context.Context) ([]*npool.PlanRelatedTestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPlanRelatedTestCaseByConds(ctx, cli); err != nil {
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

func (h *Handler) GetPlanRelatedTestCase(ctx context.Context) (info *npool.PlanRelatedTestCase, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPlanRelatedTestCase(cli); err != nil {
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
