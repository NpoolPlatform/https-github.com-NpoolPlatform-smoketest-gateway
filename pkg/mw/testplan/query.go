package testplan

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	enttestplan "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stm   *ent.TestPlanSelect
	infos []*npool.TestPlan
	total uint32
}

func (h *queryHandler) selectTestPlan(stm *ent.TestPlanQuery) {
	h.stm = stm.Select(
		enttestplan.FieldID,
		enttestplan.FieldName,
		// enttestplan.FieldState,
		enttestplan.FieldOwnerID,
		enttestplan.FieldResponsibleUserID,
		enttestplan.FieldFailedTestCasesCount,
		enttestplan.FieldPassedTestCasesCount,
		enttestplan.FieldSkippedTestCasesCount,
		enttestplan.FieldRunDuration,
		enttestplan.FieldDeadline,
		// enttestplan.FieldTestResult,
		enttestplan.FieldDeadline,
		enttestplan.FieldCreatedAt,
		enttestplan.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryJoin() {
	h.stm.Modify(func(s *sql.Selector) {})
}

func (h *queryHandler) queryTestPlan(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid testplan id")
	}
	h.selectTestPlan(
		cli.TestPlan.
			Query().
			Where(
				enttestplan.ID(uuid.MustParse(*h.ID)),
				enttestplan.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryTestPlanByConds(ctx context.Context, cli *ent.Client) (err error) {
	if h.Conds == nil {
		return fmt.Errorf("invalid conds")
	}

	stm := cli.TestPlan.Query()
	if h.Conds.ID != nil {
		stm = stm.Where(
			enttestplan.ID(uuid.MustParse(h.Conds.GetID().GetValue())),
		)
	}

	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}

	h.total = uint32(total)

	h.selectTestPlan(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stm.Scan(ctx, &h.infos)
}

func (h *Handler) GetTestPlans(ctx context.Context) ([]*npool.TestPlan, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTestPlanByConds(ctx, cli); err != nil {
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

func (h *Handler) GetTestPlan(ctx context.Context) (info *npool.TestPlan, err error) {
	handler := &queryHandler{
		Handler: h,
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryTestPlan(cli); err != nil {
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
