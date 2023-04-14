package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func CreateSet(c *ent.TestCaseCreate, in *npool.TestCaseReq) *ent.TestCaseCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	if in.ModuleID != nil {
		c.SetModuleID(uuid.MustParse(in.GetModuleID()))
	}
	if in.ApiID != nil {
		c.SetAPIID(uuid.MustParse(in.GetApiID()))
	}
	if in.Arguments != nil {
		c.SetArguments(in.GetArguments())
	}
	if in.ArgTypeDescription != nil {
		c.SetArgTypeDescription(in.GetArgTypeDescription())
	}
	if in.ExpectationResult != nil {
		c.SetExpectationResult(in.GetExpectationResult())
	}
	if in.TestCaseType != nil {
		c.SetTestCaseType(in.GetTestCaseType().String())
	}
	if in.Deprecated != nil {
		c.SetDeprecated(in.GetDeprecated())
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}
	return c
}

func Create(ctx context.Context, in *npool.TestCaseReq) (*ent.TestCase, error) {
	var info *ent.TestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.TestCase.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.TestCaseReq) ([]*ent.TestCase, error) {
	var err error

	rows := []*ent.TestCase{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		bulk := make([]*ent.TestCaseCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(cli.TestCase.Create(), info)
		}
		rows, err = cli.TestCase.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
