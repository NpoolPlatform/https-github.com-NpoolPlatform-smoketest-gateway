package relatedtestcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/relatedtestcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/relatedtestcase"

	"github.com/google/uuid"
)

func Row(ctx context.Context, testCaseID string) (*ent.RelatedTestCase, error) {
	var info *ent.RelatedTestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.RelatedTestCase.Query().Where(relatedtestcase.RelatedTestCaseID(uuid.MustParse(testCaseID))).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.RelatedTestCaseQuery, error) {
	stm := cli.RelatedTestCase.Query()
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(relatedtestcase.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid module field")
		}
	}

	if conds.ID != nil {
		switch conds.GetTestCaseID().GetOp() {
		case cruder.EQ:
			stm.Where(relatedtestcase.TestCaseID(uuid.MustParse(conds.GetTestCaseID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid module field")
		}
	}

	if conds.CondType != nil {
		switch conds.GetCondType().GetOp() {
		case cruder.EQ:
			stm.Where(relatedtestcase.CondType(conds.GetCondType().GetValue()))
		default:
			return nil, fmt.Errorf("invalid module field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.RelatedTestCase, int, error) {
	var err error

	rows := []*ent.RelatedTestCase{}
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
			Order(ent.Desc(relatedtestcase.FieldUpdatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.RelatedTestCase, error) {
	var info *ent.RelatedTestCase

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
