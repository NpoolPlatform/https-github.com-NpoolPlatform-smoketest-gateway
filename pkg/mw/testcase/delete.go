package testcase

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	testcasecrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

func (h *Handler) DeleteTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	info, err = h.GetTestCase(ctx)
	if err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		now := uint32(time.Now().Unix())
		if _, err := testcasecrud.UpdateSet(
			cli.TestCase.UpdateOneID(*h.ID),
			&testcasecrud.Req{
				ID:        h.ID,
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
