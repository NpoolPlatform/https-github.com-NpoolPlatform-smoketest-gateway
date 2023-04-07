package relatedtestcase

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/relatedtestcase"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func UpdateSet(info *ent.Module, in *npool.RelatedTestCaseReq) *ent.ModuleUpdateOne {
	return info.Update()
}

func Update(ctx context.Context, in *npool.RelatedTestCaseReq) (*ent.Module, error) {
	return nil, nil
}
