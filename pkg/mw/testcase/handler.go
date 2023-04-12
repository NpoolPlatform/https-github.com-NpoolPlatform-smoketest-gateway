package testcase

import (
	"context"
	"fmt"

	testcasemgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	"github.com/google/uuid"
)

type Handler struct {
	ID                *string
	Name              *string
	Description       *string
	ModuleID          *string
	ModuleName        *string
	APIID             *string
	Arguments         *string
	ExpectationResult *string
	TestCaseType      *testcasemgrpb.TestCaseType
	Deprecated        *bool
	Conds             *testcasemgrpb.Conds
	Offset            *int32
	Limit             *int32
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

func WithAPIID(apiID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*apiID); err != nil {
			return err
		}
		h.APIID = apiID
		return nil
	}
}

func WithModuleID(moduleID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _, err := uuid.Parse(*moduleID); err != nil {
			return err
		}
		h.ModuleID = moduleID
		return nil
	}
}

func WithModuleName(moduleName *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if moduleName == nil {
			return nil
		}
		h.ModuleName = moduleName
		return nil
	}
}

func WithExpectationResult(expectationResult *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if expectationResult == nil {
			return fmt.Errorf("invalid expectation result")
		}
		h.ExpectationResult = expectationResult
		return nil
	}
}

func WithArguments(arguments *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if arguments == nil {
			return fmt.Errorf("invalid arguments")
		}
		h.Arguments = arguments
		return nil
	}
}

func WithName(name *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if name == nil {
			return nil
		}
		if *name == "" {
			return fmt.Errorf("invalid testcase name")
		}
		h.Name = name
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

func WithConds(conds *testcasemgrpb.Conds, offset, limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return fmt.Errorf("invalid conds")
		}

		if conds.ID != nil {
			if _, err := uuid.Parse(conds.GetID().GetValue()); err != nil {
				return err
			}
		}

		h.Conds = conds
		h.Offset = &offset
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = &limit

		return nil
	}
}
