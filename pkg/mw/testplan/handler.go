package testplan

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	testplancrud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *uuid.UUID
	Name        *string
	State       *npool.TestPlanState
	CreatedBy   *uuid.UUID
	Executor    *uuid.UUID
	Deadline    *uint32
	Fails       *uint32
	Skips       *uint32
	Passes      *uint32
	Result      *npool.TestResultState
	RunDuration *uint32
	Conds       *testplancrud.Conds
	Offset      int32
	Limit       int32
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
		if id == nil {
			return nil
		}
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
		const leastNameLen = 4
		if len(*name) < leastNameLen {
			return fmt.Errorf("name %v too short", *name)
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
		if *deadline <= uint32(time.Now().Unix()) {
			return fmt.Errorf("deadline less than current time")
		}

		h.Deadline = deadline
		return nil
	}
}

func WithState(state *npool.TestPlanState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			return nil
		}
		switch *state {
		case npool.TestPlanState_WaitStart:
		case npool.TestPlanState_InProgress:
		case npool.TestPlanState_Finished:
		case npool.TestPlanState_Overdue:
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

func WithResult(result *npool.TestResultState) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			return nil
		}
		switch *result {
		case npool.TestResultState_Failed:
		case npool.TestResultState_Passed:
		default:
			return fmt.Errorf("plan result %v invalid", *result)
		}
		h.Result = result
		return nil
	}
}

func WithRunDuration(duration *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			return nil
		}
		h.RunDuration = duration
		return nil
	}
}

func WithFails(fails *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if fails == nil {
			return nil
		}
		h.Fails = fails
		return nil
	}
}

func WithPasses(passes *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if passes == nil {
			return nil
		}
		h.Passes = passes
		return nil
	}
}

func WithSkips(skips *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if skips == nil {
			return nil
		}
		h.Skips = skips
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
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
