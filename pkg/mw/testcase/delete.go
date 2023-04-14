package testcase

import (
	"context"
	"time"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteTestCase(ctx context.Context, tx *ent.Tx) error {
	if _, err := tx.
		TestCase.
		UpdateOneID(uuid.MustParse(*h.ID)).
		SetDeletedAt(uint32(time.Now().Unix())).
		Save(ctx); err != nil {
		if !ent.IsNotFound(err) {
			return err
		}
	}
	return nil
}

func (h *Handler) DeleteTestCase(ctx context.Context) (info *npool.TestCase, err error) {
	info, err = h.GetTestCase(ctx)
	if err != nil {
		return nil, err
	}

	handler := &deleteHandler{
		Handler: h,
	}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteTestCase(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return info, nil
}
