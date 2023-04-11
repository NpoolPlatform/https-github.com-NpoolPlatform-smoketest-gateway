package planrelatedtestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/planrelatedtestcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func CreateSet(c *ent.PlanRelatedTestCaseCreate, in *npool.PlanRelatedTestCaseReq) *ent.PlanRelatedTestCaseCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.TestCaseID != nil {
		c.SetTestCaseID(uuid.MustParse(in.GetTestCaseID()))
	}
	if in.TestPlanID != nil {
		c.SetTestPlanID(uuid.MustParse(in.GetTestPlanID()))
	}
	if in.TestUserID != nil {
		c.SetTestUserID(uuid.MustParse(in.GetTestUserID()))
	}
	if in.TestCaseOutput != nil {
		c.SetTestCaseOutput(in.GetTestCaseOutput())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	if in.RunDuration != nil {
		c.SetRunDuration(in.GetRunDuration())
	}
	if in.TestCaseResult != nil {
		c.SetTestCaseResult(in.GetTestCaseResult().Enum().String())
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}
	return c
}

func Create(ctx context.Context, in *npool.PlanRelatedTestCaseReq) (*ent.PlanRelatedTestCase, error) {
	var info *ent.PlanRelatedTestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.PlanRelatedTestCase.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.PlanRelatedTestCaseReq) ([]*ent.PlanRelatedTestCase, error) {
	var err error

	rows := []*ent.PlanRelatedTestCase{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		bulk := make([]*ent.PlanRelatedTestCaseCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(cli.PlanRelatedTestCase.Create(), info)
		}
		rows, err = cli.PlanRelatedTestCase.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
