package cond

import (
	"context"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	return nil
}

func (h *Handler) CreateCond(ctx context.Context) (info *npool.Cond, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := crud.CreateSet(
			cli.Cond.Create(),
			&crud.Req{
				TestCaseID:     h.TestCaseID,
				CondTestCaseID: h.CondTestCaseID,
				CondType:       h.CondType,
				ArgumentMap:    h.ArgumentMap,
				Index:          h.Index,
			},
		).Save(ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})
	if err != nil {
		return nil, err
	}
	return h.GetCond(ctx)
}
