package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"

	"github.com/google/uuid"
)

func Row(ctx context.Context, id string) (*ent.TestCase, error) {
	var info *ent.TestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.TestCase.Query().Where(testcase.ID(uuid.MustParse(id))).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.TestCaseQuery, error) { //nolint
	stm := cli.TestCase.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(testcase.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid module field")
		}
	}

	if conds.ModuleID != nil {
		switch conds.GetModuleID().GetOp() {
		case cruder.EQ:
			stm.Where(testcase.ModuleID(uuid.MustParse(conds.GetModuleID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid module field")
		}
	}

	if conds.Deprecated != nil {
		switch conds.GetDeprecated().GetOp() {
		case cruder.EQ:
			stm.Where(testcase.Deprecated(conds.GetDeprecated().GetValue()))
		default:
			return nil, fmt.Errorf("invalid module field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.TestCase, int, error) {
	var err error

	rows := []*ent.TestCase{}
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
			Order(ent.Desc(testcase.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.TestCase, error) {
	var info *ent.TestCase

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
