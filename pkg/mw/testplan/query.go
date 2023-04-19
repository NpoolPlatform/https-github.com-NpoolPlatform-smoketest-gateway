package testplan

import (
	"context"
	"fmt"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	enttestplan "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
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
		enttestplan.FieldState,
		enttestplan.FieldCreatedBy,
		enttestplan.FieldExecutor,
		enttestplan.FieldPasses,
		enttestplan.FieldSkips,
		enttestplan.FieldFails,
		enttestplan.FieldRunDuration,
		enttestplan.FieldDeadline,
		enttestplan.FieldResult,
		enttestplan.FieldDeadline,
		enttestplan.FieldCreatedAt,
		enttestplan.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryTestPlan(cli *ent.Client) error {
	if h.ID == nil {
		return fmt.Errorf("invalid testplan id")
	}
	h.selectTestPlan(
		cli.TestPlan.
			Query().
			Where(
				enttestplan.ID(*h.ID),
				enttestplan.DeletedAt(0),
			),
	)
	return nil
}

func (h *queryHandler) queryTestPlanByConds(ctx context.Context, cli *ent.Client) (err error) {
	stm, err := testplancrud.SetQueryConds(cli.TestPlan.Query(), h.Conds)
	if err != nil {
		return err
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
