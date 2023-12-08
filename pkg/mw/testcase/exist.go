package testcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	enttestcase "github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
)

func (h *Handler) ExistTestCase(ctx context.Context) (exist bool, err error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			TestCase.
			Query().
			Where(enttestcase.EntID(*h.EntID), enttestcase.DeletedAt(0)).
			Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	if !exist && err == nil {
		return exist, fmt.Errorf("id %v not exist", *h.EntID)
	}

	return exist, nil
}
