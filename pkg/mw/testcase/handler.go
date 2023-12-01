package testcase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	module1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/module"
	"github.com/google/uuid"
)

type Handler struct {
	ID            *uint32
	EntID         *uuid.UUID
	Name          *string
	Description   *string
	ModuleID      *uuid.UUID
	ApiID         *uuid.UUID //nolint
	Input         *string
	InputDesc     *string
	Expectation   *string
	OutputDesc    *string
	TestCaseType  *npool.TestCaseType
	TestCaseClass *npool.TestCaseClass
	Deprecated    *bool
	Conds         *crud.Conds
	Offset        int32
	Limit         int32
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
		const leastNameLen = 2
		if len(*name) < leastNameLen {
			return fmt.Errorf("name %v too short", *name)
		}
		h.Name = name
		return nil
	}
}

func WithDescription(description *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Description = description
		return nil
	}
}

func WithModuleID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := module1.NewHandler(
			ctx,
			module1.WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistModule(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid module")
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ModuleID = &_id
		return nil
	}
}

func WithExpectation(expectation *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if expectation == nil {
			if must {
				return fmt.Errorf("invalid expectation")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*expectation), &r); err != nil {
			return err
		}
		h.Expectation = expectation
		return nil
	}
}

func WithOutputDesc(outputDesc *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if outputDesc == nil {
			if must {
				return fmt.Errorf("invalid outputdesc")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*outputDesc), &r); err != nil {
			return err
		}
		h.OutputDesc = outputDesc
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

func WithInputDesc(inputDesc *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if inputDesc == nil {
			if must {
				return fmt.Errorf("invalid inputdesc")
			}
			return nil
		}

		var r interface{}
		if err := json.Unmarshal([]byte(*inputDesc), &r); err != nil {
			return err
		}
		h.InputDesc = inputDesc
		return nil
	}
}

func WithDeprecated(deprecated *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Deprecated = deprecated
		return nil
	}
}

func WithTestCaseType(testCaseType *npool.TestCaseType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseType == nil {
			if must {
				return fmt.Errorf("invalid testcasetype")
			}
			return nil
		}
		switch *testCaseType {
		case npool.TestCaseType_Manual:
		case npool.TestCaseType_Automatic:
		default:
			return fmt.Errorf("invalid testcase type")
		}
		h.TestCaseType = testCaseType
		return nil
	}
}

func WithTestCaseClass(testCaseClass *npool.TestCaseClass, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseClass == nil {
			if must {
				return fmt.Errorf("invalid testcaseclass")
			}
			return nil
		}
		switch *testCaseClass {
		case npool.TestCaseClass_Functionality:
		case npool.TestCaseClass_Interface:
		default:
			return fmt.Errorf("invalid testcase class")
		}
		h.TestCaseClass = testCaseClass
		return nil
	}
}

//nolint
func WithApiID(apiID *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if apiID == nil {
			if must {
				return fmt.Errorf("invalid apiid")
			}
			return nil
		}
		_apiID, err := uuid.Parse(*apiID)
		if err != nil {
			return err
		}
		h.ApiID = &_apiID
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

		if conds.ModuleID != nil {
			id, err := uuid.Parse(conds.GetModuleID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ModuleID = &cruder.Cond{Op: conds.ModuleID.Op, Val: id}
		}
		if conds.ApiID != nil {
			id, err := uuid.Parse(conds.GetApiID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ApiID = &cruder.Cond{Op: conds.ApiID.Op, Val: id}
		}
		if conds.Deprecated != nil {
			h.Conds.Deprecated = &cruder.Cond{Op: conds.GetDeprecated().GetOp(), Val: conds.GetDeprecated().GetValue()}
		}

		if len(conds.GetEntIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}
		if len(conds.GetIDs().GetValue()) > 0 {
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: conds.GetIDs().GetValue()}
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
