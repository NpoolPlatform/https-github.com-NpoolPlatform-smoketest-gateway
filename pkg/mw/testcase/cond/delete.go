package cond

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) DeleteCond(ctx context.Context) (*npool.Cond, error) {
	info, err := h.GetCond(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := crud.UpdateSet(
			cli.Cond.UpdateOneID(*h.ID),
			&crud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
