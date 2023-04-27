package testcase

import (
	"context"
	"encoding/json"
	"fmt"

	apimwcli "github.com/NpoolPlatform/basal-middleware/pkg/client/api"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase"
	"github.com/google/uuid"
)

type Handler struct {
	ID           *uuid.UUID
	Name         *string
	Description  *string
	ModuleID     *uuid.UUID
	ModuleName   *string
	ApiID        *uuid.UUID //nolint
	Input        *string
	InputDesc    *string
	Expectation  *string
	TestCaseType *npool.TestCaseType
	Deprecated   *bool
	Conds        *crud.Conds
	Offset       int32
	Limit        int32
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

func WithDescription(description *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if description == nil {
			return nil
		}
		h.Description = description
		return nil
	}
}

func WithExpectation(expectation *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if expectation == nil {
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

func WithInput(input *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if input == nil {
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

func WithInputDesc(inputDesc *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if inputDesc == nil {
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

func WithDeprecated(deprecated *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if deprecated == nil {
			return nil
		}
		h.Deprecated = deprecated
		return nil
	}
}

func WithTestCaseType(testCaseType *npool.TestCaseType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if testCaseType == nil {
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

//nolint
func WithApiID(apiID *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_apiID, err := uuid.Parse(*apiID)
		if err != nil {
			return err
		}

		exist, err := apimwcli.ExistAPI(ctx, *apiID)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid api id")
		}

		h.ApiID = &_apiID
		return nil
	}
}

func WithModuleName(moduleName *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if moduleName == nil {
			return fmt.Errorf("invalid module name")
		}
		const leastNameLen = 2
		if len(*moduleName) < leastNameLen {
			return fmt.Errorf("name %v too short", *moduleName)
		}
		h.ModuleName = moduleName
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
			h.Conds.ID = &cruder.Cond{Op: conds.ID.Op, Val: id}
		}

		if conds.ModuleID != nil {
			id, err := uuid.Parse(conds.GetModuleID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ModuleID = &cruder.Cond{Op: conds.ModuleID.Op, Val: id}
		}

		if conds.Deprecated != nil {
			h.Conds.Deprecated = &cruder.Cond{Op: conds.Deprecated.Op, Val: conds.Deprecated}
		}

		if len(conds.GetIDs().GetValue()) > 0 {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: ids}
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
