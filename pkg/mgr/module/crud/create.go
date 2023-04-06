package detail

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/module"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func CreateSet(c *ent.ModuleCreate, in *npool.ModuleReq) *ent.ModuleCreate {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.Name != nil {
		c.SetName(in.GetName())
	}
	if in.Description != nil {
		c.SetDescription(in.GetDescription())
	}
	if in.CreatedAt != nil {
		c.SetCreatedAt(in.GetCreatedAt())
	}
	return c
}

func Create(ctx context.Context, in *npool.ModuleReq) (*ent.Module, error) {
	var info *ent.Module
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.Module.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
