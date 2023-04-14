package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func CreateSet(c *ent.TestPlanCreate, in *npool.TestPlanReq) *ent.TestPlanCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.State != nil {
		c.SetState(in.GetState().String())
	}
	if in.OwnerID != nil {
		c.SetOwnerID(uuid.MustParse(in.GetOwnerID()))
	}
	if in.ResponsibleUserID != nil {
		c.SetResponsibleUserID(uuid.MustParse(in.GetResponsibleUserID()))
	}
	if in.FailedTestCasesCount != nil {
		c.SetFailedTestCasesCount(in.GetFailedTestCasesCount())
	}
	if in.PassedTestCasesCount != nil {
		c.SetPassedTestCasesCount(in.GetPassedTestCasesCount())
	}
	if in.SkippedTestCasesCount != nil {
		c.SetSkippedTestCasesCount(in.GetSkippedTestCasesCount())
	}
	if in.RunDuration != nil {
		c.SetRunDuration(in.GetRunDuration())
	}
	if in.TestResult != nil {
		c.SetTestResult(in.GetTestResult().String())
	}
	if in.Deadline != nil {
		c.SetDeadline(in.GetDeadline())
	}
	if in.CreatedAt != nil {
		c.SetDeadline(in.GetDeadline())
	}
	return c
}

func Create(ctx context.Context, in *npool.TestPlanReq) (*ent.TestPlan, error) {
	var info *ent.TestPlan
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.TestPlan.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.TestPlanReq) ([]*ent.TestPlan, error) {
	var err error

	rows := []*ent.TestPlan{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		bulk := make([]*ent.TestPlanCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(cli.TestPlan.Create(), info)
		}
		rows, err = cli.TestPlan.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
