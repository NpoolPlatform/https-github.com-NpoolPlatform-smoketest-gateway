package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
)

func Exist(ctx context.Context, id string) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.TestCase.Query().Where(testcase.ID(uuid.MustParse(id))).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
