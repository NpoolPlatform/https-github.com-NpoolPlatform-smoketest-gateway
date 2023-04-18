package plantestcase

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	plantestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.PlanTestCaseSelect
	infos []*npool.PlanTestCase
	total uint32
}

func (h *queryHandler) selectPlanTestCase(stm *ent.PlanTestCaseQuery) {
	h.stm = stm.Select(
		plantestcase.FieldID,
		plantestcase.FieldTestPlanID,
		plantestcase.FieldTestCaseID,
		plantestcase.FieldTestCaseOutput,
		// plantestcase.FieldTestCaseResult,
		plantestcase.FieldTestUserID,
		plantestcase.FieldRunDuration,
		plantestcase.FieldDescription,
		plantestcase.FieldCreatedAt,
		plantestcase.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) queryPlanTestCase(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid plantestcase id")
	}
	h.selectPlanTestCase(
		cli.PlanTestCase.
			Query().
			Where(
				plantestcase.ID(uuid.MustParse(*h.ID)),
				plantestcase.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryPlanTestCaseByConds(ctx context.Context, cli *ent.Client) (err error) {
	if h.Conds == nil {
		return fmt.Errorf("invalid conds")
	}

	stm := cli.PlanTestCase.Query()
	if h.Conds.ID != nil {
		stm = stm.Where(
			plantestcase.ID(uuid.MustParse(h.Conds.GetID().GetValue())),
		)
	}
	if h.Conds.TestPlanID != nil {
		stm = stm.Where(
			plantestcase.TestPlanID(uuid.MustParse(h.Conds.GetTestPlanID().GetValue())),
		)
	}
	if h.Conds.TestUserID != nil {
		stm = stm.Where(
			plantestcase.TestUserID(uuid.MustParse(h.Conds.GetTestPlanID().GetValue())),
		)
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectPlanTestCase(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetPlanTestCases(ctx context.Context) ([]*npool.PlanTestCase, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPlanTestCaseByConds(ctx, cli); err != nil {
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

func (h *Handler) GetPlanTestCase(ctx context.Context) (info *npool.PlanTestCase, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPlanTestCase(cli); err != nil {
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
