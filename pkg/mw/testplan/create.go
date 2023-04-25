package testplan

import (
	"context"
	"fmt"
	"time"

	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) validate() error {
	if h.Name == nil {
		return fmt.Errorf("invalid name")
	}
	const leastNameLen = 4
	if len(*h.Name) < leastNameLen {
		return fmt.Errorf("name %v too short", *h.Name)
	}

	if h.Deadline != nil {
		if *h.Deadline <= uint32(time.Now().Unix()) {
			return fmt.Errorf("deadline less than current time")
		}
	}
	return nil
}

func (h *Handler) CreateTestPlan(ctx context.Context) (info *npool.TestPlan, err error) {
	handler := &createHandler{
		Handler: h,
	}

	if err := handler.validate(); err != nil {
		return nil, err
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := testplancrud.CreateSet(
			cli.TestPlan.Create(),
			&testplancrud.Req{
				Name:      h.Name,
				CreatedBy: h.CreatedBy,
				Executor:  h.Executor,
				Deadline:  h.Deadline,
			},
		).Save(_ctx)
		if err != nil {
			return err
		}

		h.ID = &info.ID
		return nil
	})

	if err != nil {
		return nil, err
	}

	return h.GetTestPlan(ctx)
}
