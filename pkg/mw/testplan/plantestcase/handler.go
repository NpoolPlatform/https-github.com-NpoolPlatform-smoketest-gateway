package plantestcase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	testplan1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"
	"github.com/google/uuid"
)

type Handler struct {
	ID          *uint32
	EntID       *uuid.UUID
	TestPlanID  *uuid.UUID
	TestCaseID  *uuid.UUID
	TestUserID  *uuid.UUID
	Input       *string
	Output      *string
	Result      *npool.TestCaseResult
	Description *string
	Index       *uint32
	RunDuration *uint32
	Conds       *crud.Conds
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

//nolint
func WithTestPlanID(planID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := testplan1.NewHandler(
			ctx,
			testplan1.WithEntID(planID, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistTestPlan(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid testplan")
		}
		_planID, err := uuid.Parse(*planID)
		if err != nil {
			return err
		}
		h.TestPlanID = &_planID
		return nil
	}
}

//nolint
func WithTestCaseID(testCaseID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := testcase1.NewHandler(
			ctx,
			testcase1.WithEntID(testCaseID, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistTestCase(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid testcase")
		}
		_testCaseID, err := uuid.Parse(*testCaseID)
		if err != nil {
			return err
		}
		h.TestCaseID = &_testCaseID
		return nil
	}
}

func WithTestUserID(userID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
			if must {
				return fmt.Errorf("invalid testuserid")
			}
			return nil
		}
		_userID, err := uuid.Parse(*userID)
		if err != nil {
			return err
		}

		h.TestUserID = &_userID
		return nil
	}
}

func WithInput(input *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if input == nil {
			if must {
				return fmt.Errorf("invalid input")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*input), &r); err != nil {
			return err
		}
		h.Input = input
		return nil
	}
}

func WithOutput(output *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if output == nil {
			if must {
				return fmt.Errorf("invalid output")
			}
			return nil
		}
		h.Output = output
		return nil
	}
}

func WithResult(result *npool.TestCaseResult, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			if must {
				return fmt.Errorf("invalid result")
			}
			return nil
		}
		switch *result {
		case npool.TestCaseResult_Passed:
		case npool.TestCaseResult_Failed:
		case npool.TestCaseResult_Skipped:
		default:
			return fmt.Errorf("invalid testcase result")
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

func WithIndex(index *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = index
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Description = description
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
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

		if conds.TestPlanID != nil {
			id, err := uuid.Parse(conds.GetTestPlanID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TestPlanID = &cruder.Cond{Op: conds.TestPlanID.Op, Val: id}
		}

		if conds.TestUserID != nil {
			id, err := uuid.Parse(conds.GetTestUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TestUserID = &cruder.Cond{Op: conds.TestUserID.Op, Val: id}
		}

		if conds.Result != nil {
			switch conds.GetResult().GetValue() {
			case uint32(npool.TestCaseResult_Passed):
			case uint32(npool.TestCaseResult_Failed):
			case uint32(npool.TestCaseResult_Skipped):
			default:
				return fmt.Errorf("invalid testplanstate")
			}
			_state := conds.GetResult().GetValue()
			h.Conds.Result = &cruder.Cond{Op: conds.GetResult().GetOp(), Val: npool.TestCaseResult(_state)}
		}

		if len(conds.GetTestPlanIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetTestPlanIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.TestPlanIDs = &cruder.Cond{Op: conds.GetTestPlanIDs().GetOp(), Val: ids}
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
