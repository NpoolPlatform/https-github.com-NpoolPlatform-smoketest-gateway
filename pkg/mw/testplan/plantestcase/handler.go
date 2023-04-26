package plantestcase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testplan/plantestcase"
	testcasemw "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	testplanmw "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testplan"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	TestPlanID     *uuid.UUID
	TestCaseID     *uuid.UUID
	TestUserID     *uuid.UUID
	TestCaseOutput *string
	TestCaseResult *npool.TestCaseResult
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

//nolint
func WithTestPlanID(planID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_planID, err := uuid.Parse(*planID)
		if err != nil {
			return err
		}

		type TestPlanHandler struct {
			testplanmw.Handler
		}

		testplan := &TestPlanHandler{}
		testplan.ID = &_planID

		if _, err := testplan.ExistTestPlan(ctx); err != nil {
			return err
		}

		h.TestPlanID = &_planID
		return nil
	}
}

//nolint
func WithTestCaseID(testCaseID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_testCaseID, err := uuid.Parse(*testCaseID)
		if err != nil {
			return err
		}

		type TestCaseHandler struct {
			testcasemw.Handler
		}

		testcase := &TestCaseHandler{}
		testcase.ID = &_testCaseID

		if _, err := testcase.ExistTestCase(ctx); err != nil {
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

func WithTestCaseResult(result *npool.TestCaseResult) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if result == nil {
			return nil
		}
		switch *result {
		case npool.TestCaseResult_Passed:
		case npool.TestCaseResult_Failed:
		case npool.TestCaseResult_Skipped:
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
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

		if len(conds.GetTestPlanIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetTestPlanIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.TestPlanID = &cruder.Cond{Op: conds.GetTestPlanIDs().GetOp(), Val: ids}
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
