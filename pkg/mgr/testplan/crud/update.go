package testplan

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func UpdateSet(info *ent.TestPlan, in *npool.TestPlanReq) *ent.TestPlanUpdateOne {
	return info.Update()
}

func Update(ctx context.Context, in *npool.TestPlanReq) (*ent.TestPlan, error) {
	return nil, nil
}
