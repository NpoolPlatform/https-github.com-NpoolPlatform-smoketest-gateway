package testcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func UpdateSet(info *ent.Module, in *npool.TestCaseReq) *ent.ModuleUpdateOne {
	return info.Update()
}

func Update(ctx context.Context, in *npool.TestCaseReq) (*ent.Module, error) {
	return nil, nil
}
