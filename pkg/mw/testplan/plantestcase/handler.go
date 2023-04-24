package plantestcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/plantestcase"
	testcasecli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testcase"
	testplancli "github.com/NpoolPlatform/smoketest-middleware/pkg/client/testplan"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	TestPlanID     *uuid.UUID
	TestCaseID     *uuid.UUID
	TestUserID     *uuid.UUID
	TestCaseOutput *string
	TestCaseResult *mgrpb.TestCaseResult
	Description    *string
	Index          *uint32
	RunDuration    *uint32
	Conds          *crud.Conds
	Offset         int32
	Limit          int32
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

func WithTestPlanID(planID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_planID, err := uuid.Parse(*planID)
		if err != nil {
			return err
		}

		if exist, err := testplancli.ExistTestPlan(ctx, *planID); !exist {
			return err
		}

		h.TestPlanID = &_planID
		return nil
	}
}

func WithTestCaseID(testCaseID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_testCaseID, err := uuid.Parse(*testCaseID)
		if err != nil {
			return err
		}

		if _, err := testcasecli.ExistTestCase(ctx, *testCaseID); err != nil {
			return err
		}

		h.TestCaseID = &_testCaseID
		return nil
	}
}

func WithTestUserID(userID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if userID == nil {
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

func WithTestCaseOutput(output *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if output == nil {
			return nil
		}
		h.TestCaseOutput = output
		return nil
	}
}

func WithTestCaseResult(result *mgrpb.TestCaseResult) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			return fmt.Errorf("need testcase result")
		}
		switch *result {
		case mgrpb.TestCaseResult_Passed:
		case mgrpb.TestCaseResult_Failed:
		case mgrpb.TestCaseResult_Skipped:
		default:
			return fmt.Errorf("invalid testcase result")
		}
		h.TestCaseResult = result
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

func WithIndex(index *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if index == nil {
			return nil
		}
		h.Index = index
		return nil
	}
}

func WithDescription(description *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			return nil
		}
		h.Description = description
		return nil
	}
}

func WithConds(conds *mgrpb.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &crud.Conds{}
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

		if conds.TestPlanID != nil {
			id, err := uuid.Parse(conds.GetTestPlanID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TestPlanID = &cruder.Cond{Op: h.Conds.TestPlanID.Op, Val: id}
		}

		if conds.TestUserID != nil {
			id, err := uuid.Parse(conds.GetTestUserID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TestUserID = &cruder.Cond{Op: h.Conds.TestUserID.Op, Val: id}
		}

		if conds.Result != nil {
			h.Conds.Result = &cruder.Cond{Op: h.Conds.Result.Op, Val: conds.Result}
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
