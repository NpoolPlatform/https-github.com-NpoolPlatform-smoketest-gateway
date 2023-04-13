package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func UpdateSet(info *ent.TestCaseUpdateOne, in *npool.TestCaseReq) *ent.TestCaseUpdateOne {
	if in.Name != nil {
		info.SetName(in.GetName())
	}
	if in.Description != nil {
		info.SetDescription(in.GetDescription())
	}
	if in.Arguments != nil {
		info.SetArguments(in.GetArguments())
	}
	if in.ExpectationResult != nil {
		info.SetExpectationResult(in.GetExpectationResult())
	}
	if in.TestCaseType != nil {
		info.SetTestCaseType(in.GetTestCaseType().String())
	}
	if in.Deprecated != nil {
		info.SetDeprecated(in.GetDeprecated())
	}
	return info
}

func Update(ctx context.Context, in *npool.TestCaseReq) (*ent.TestCaseUpdateOne, error) {
	return nil, nil
}
