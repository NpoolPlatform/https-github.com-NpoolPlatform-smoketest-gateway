package testplan

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	enttestplan "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
)

func (h *Handler) ExistTestPlan(ctx context.Context) (exist bool, err error) {
	if h.ID == nil {
		return false, fmt.Errorf("invalid id")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			TestPlan.
			Query().
			Where(enttestplan.ID(*h.ID), enttestplan.DeletedAt(0)).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}
