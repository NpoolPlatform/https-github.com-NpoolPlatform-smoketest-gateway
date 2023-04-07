package planrelatedtestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/planrelatedtestcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func UpdateSet(info *ent.PlanRelatedTestCase, in *npool.PlanRelatedTestCaseReq) *ent.PlanRelatedTestCaseUpdateOne {
	return info.Update()
}

func Update(ctx context.Context, in *npool.PlanRelatedTestCaseReq) (*ent.PlanRelatedTestCase, error) {
	return nil, nil
}
