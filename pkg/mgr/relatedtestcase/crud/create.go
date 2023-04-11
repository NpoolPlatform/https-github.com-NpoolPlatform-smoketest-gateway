package relatedtestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/relatedtestcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func CreateSet(c *ent.RelatedTestCaseCreate, in *npool.RelatedTestCaseReq) *ent.RelatedTestCaseCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.TestCaseID != nil {
		c.SetTestCaseID(uuid.MustParse(in.GetTestCaseID()))
	}
	if in.RelatedTestCaseID != nil {
		c.SetRelatedTestCaseID(uuid.MustParse(in.GetRelatedTestCaseID()))
	}
	if in.ArgumentsTransfer != nil {
		c.SetArgumentsTransfer(in.GetArgumentsTransfer())
	}
	if in.Index != nil {
		c.SetIndex(in.GetIndex())
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}
	return c
}

func Create(ctx context.Context, in *npool.RelatedTestCaseReq) (*ent.RelatedTestCase, error) {
	var info *ent.RelatedTestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.RelatedTestCase.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.RelatedTestCaseReq) ([]*ent.RelatedTestCase, error) {
	var err error

	rows := []*ent.RelatedTestCase{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		bulk := make([]*ent.RelatedTestCaseCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(cli.RelatedTestCase.Create(), info)
		}
		rows, err = cli.RelatedTestCase.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
