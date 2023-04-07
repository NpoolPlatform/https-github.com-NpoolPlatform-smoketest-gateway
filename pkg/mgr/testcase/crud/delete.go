package testcase

import (
	"context"
	"time"

	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

func Delete(ctx context.Context, id string) (*ent.TestCase, error) {
	var info *ent.TestCase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.TestCase.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
