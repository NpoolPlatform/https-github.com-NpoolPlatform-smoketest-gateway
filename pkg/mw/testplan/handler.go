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
	ID          *uint32
	EntID       *uuid.UUID
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

func WithID(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = u
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithName(name *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
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

func WithCreatedBy(createdBy *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if createdBy == nil {
			if must {
				return fmt.Errorf("invalid createdby")
			}
			return nil
		}
		_createdBy, err := uuid.Parse(*createdBy)
		if err != nil {
			return err
		}
		h.CreatedBy = &_createdBy
		return nil
	}
}

func WithDeadline(deadline *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deadline == nil {
			if must {
				return fmt.Errorf("invalid deadline")
			}
			return nil
		}
		if *deadline <= uint32(time.Now().Unix()) {
			return fmt.Errorf("deadline less than current time")
		}

		h.Deadline = deadline
		return nil
	}
}

func WithState(state *npool.TestPlanState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if state == nil {
			if must {
				return fmt.Errorf("invalid state")
			}
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

func WithExecutor(executor *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if executor == nil {
			if must {
				return fmt.Errorf("invalid executor")
			}
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

func WithResult(result *npool.TestResultState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			if must {
				return fmt.Errorf("invalid result")
			}
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

func WithRunDuration(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RunDuration = duration
		return nil
	}
}

func WithFails(fails *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Fails = fails
		return nil
	}
}

func WithPasses(passes *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Passes = passes
		return nil
	}
}

func WithSkips(skips *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
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
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}

		if conds.State != nil {
			switch conds.GetState().GetValue() {
			case uint32(npool.TestPlanState_WaitStart):
			case uint32(npool.TestPlanState_InProgress):
			case uint32(npool.TestPlanState_Finished):
			case uint32(npool.TestPlanState_Overdue):
			default:
				return fmt.Errorf("invalid testplanstate")
			}
			_state := conds.GetState().GetValue()
			h.Conds.State = &cruder.Cond{Op: conds.GetState().GetOp(), Val: npool.TestPlanState(_state)}
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
