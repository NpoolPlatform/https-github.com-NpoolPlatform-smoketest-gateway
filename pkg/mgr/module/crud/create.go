package module

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

func CreateBulk(ctx context.Context, in []*npool.ModuleReq) ([]*ent.Module, error) {
	var err error

	rows := []*ent.Module{}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		bulk := make([]*ent.ModuleCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(cli.Module.Create(), info)
		}
		rows, err = cli.Module.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}
