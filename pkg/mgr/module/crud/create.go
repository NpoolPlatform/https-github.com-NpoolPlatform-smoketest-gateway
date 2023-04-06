package detail

import (
	"context"

	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/message/const"

	tracer "github.com/NpoolPlatform/smoketest-middleware/pkg/mgr/module/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

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

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

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
