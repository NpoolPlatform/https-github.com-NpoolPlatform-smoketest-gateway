package module

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func UpdateSet(info *ent.Module, in *npool.ModuleReq) *ent.ModuleUpdateOne {
	return info.Update()
}

func Update(ctx context.Context, in *npool.ModuleReq) (*ent.Module, error) {
	return nil, nil
}
