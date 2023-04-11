package planrelatedtestcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/planrelatedtestcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/planrelatedtestcase"

	"github.com/google/uuid"
)

func Row(ctx context.Context, id string) (*ent.PlanRelatedTestCase, error) {
	var info *ent.PlanRelatedTestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.PlanRelatedTestCase.Query().Where(planrelatedtestcase.ID(uuid.MustParse(id))).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.PlanRelatedTestCaseQuery, error) {
	stm := cli.PlanRelatedTestCase.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(planrelatedtestcase.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid planrelatedtestcase field")
		}
	}

	if conds.TestPlanID != nil {
		switch conds.GetTestPlanID().GetOp() {
		case cruder.EQ:
			stm.Where(planrelatedtestcase.TestPlanID(uuid.MustParse(conds.GetTestPlanID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid planrelatedtestcase field")
		}
	}

	if conds.TestUserID != nil {
		switch conds.GetTestUserID().GetOp() {
		case cruder.EQ:
			stm.Where(planrelatedtestcase.TestUserID(uuid.MustParse(conds.GetTestUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid planrelatedtestcase field")
		}
	}

	if conds.TestCaseResult != nil {
		switch conds.GetTestCaseResult().GetOp() {
		case cruder.EQ:
			stm.Where(planrelatedtestcase.TestCaseResult(conds.GetTestCaseResult().GetValue()))
		default:
			return nil, fmt.Errorf("invalid planrelatedtestcase field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.PlanRelatedTestCase, int, error) {
	var err error

	rows := []*ent.PlanRelatedTestCase{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(planrelatedtestcase.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.PlanRelatedTestCase, error) {
	var info *ent.PlanRelatedTestCase

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
