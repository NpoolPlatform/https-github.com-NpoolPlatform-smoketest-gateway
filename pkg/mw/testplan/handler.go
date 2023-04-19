package testplan

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	testplanmgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/google/uuid"
)

type Handler struct {
	ID        *uuid.UUID
	Name      *string
	State     *testplanmgrpb.TestPlanState
	CreatedBy *uuid.UUID
	Executor  *uuid.UUID
	Deadline  *uint32
	Conds     *testplancrud.Conds
	Offset    int32
	Limit     int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		h.Name = name
		return nil
	}
}

func WithCreatedBy(createdBy *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_createdBy, err := uuid.Parse(*createdBy)
		if err != nil {
			return err
		}
		h.CreatedBy = &_createdBy
		return nil
	}
}

func WithDeadline(deadline *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deadline == nil {
			return nil
		}
		h.Deadline = deadline
		return nil
	}
}

func WithState(state *testplanmgrpb.TestPlanState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case testplanmgrpb.TestPlanState_WaitStart:
		case testplanmgrpb.TestPlanState_InProgress:
		case testplanmgrpb.TestPlanState_Finished:
		case testplanmgrpb.TestPlanState_Overdue:
		default:
			return fmt.Errorf("plan state %v invalid", *state)
		}
		h.State = state
		return nil
	}
}

func WithExecutor(executor *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if executor == nil {
			return nil
		}
		_executor, err := uuid.Parse(*executor)
		if err != nil {
			return err
		}

		h.Executor = &_executor
		return nil
	}
}

func WithConds(conds *testplanmgrpb.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &testplancrud.Conds{}
		if conds == nil {
			return nil
		}

		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: h.Conds.ID.Op, Val: id}
		}

		if conds.State != nil {
			h.Conds.State = &cruder.Cond{Op: h.Conds.State.Op, Val: conds.State}
		}
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
