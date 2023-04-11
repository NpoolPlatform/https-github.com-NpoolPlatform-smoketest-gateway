package testplan

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"

	"github.com/google/uuid"
)

func Row(ctx context.Context, id string) (*ent.TestPlan, error) {
	var info *ent.TestPlan
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.TestPlan.Query().Where(testplan.ID(uuid.MustParse(id))).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.TestPlanQuery, error) {
	stm := cli.TestPlan.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(testplan.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid testplan field")
		}
	}

	if conds.ResponsibleUserID != nil {
		switch conds.GetResponsibleUserID().GetOp() {
		case cruder.EQ:
			stm.Where(testplan.ResponsibleUserID(uuid.MustParse(conds.GetResponsibleUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid testplan field")
		}
	}

	if conds.OwnerID != nil {
		switch conds.GetOwnerID().GetOp() {
		case cruder.EQ:
			stm.Where(testplan.OwnerID(uuid.MustParse(conds.GetOwnerID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid testplan field")
		}
	}

	if conds.State != nil {
		switch conds.GetState().GetOp() {
		case cruder.EQ:
			stm.Where(testplan.State(conds.GetState().GetValue()))
		default:
			return nil, fmt.Errorf("invalid testplan field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.TestPlan, int, error) {
	var err error

	rows := []*ent.TestPlan{}
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
			Order(ent.Desc(testplan.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.TestPlan, error) {
	var info *ent.TestPlan

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
