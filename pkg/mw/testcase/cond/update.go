package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) UpdateCond(ctx context.Context) (info *npool.Cond, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := crud.UpdateSet(
			cli.Cond.UpdateOneID(*h.ID),
			&crud.Req{
				Index:       h.Index,
				ArgumentMap: h.ArgumentMap,
				CondType:    h.CondType,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	info, err = h.GetCond(ctx)
	if err != nil {
		return nil, err
	}

	return info, nil
}
