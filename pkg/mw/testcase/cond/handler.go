package cond

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	testcase1 "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	TestCaseID     *uuid.UUID
	CondTestCaseID *uuid.UUID
	CondType       *npool.CondType
	Index          *uint32
	ArgumentMap    *string
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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
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
func WithTestCaseID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := testcase1.NewHandler(
			ctx,
			testcase1.WithID(id, true),
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.TestCaseID = &_id
		return nil
	}
}

//nolint
func WithCondTestCaseID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := testcase1.NewHandler(
			ctx,
			testcase1.WithID(id, true),
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
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CondTestCaseID = &_id
		return nil
	}
}

func WithCondType(_type *npool.CondType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			if must {
				return fmt.Errorf("invalid type")
			}
			return nil
		}
		switch *_type {
		case npool.CondType_PreCondition:
		case npool.CondType_Cleaner:
		default:
			return fmt.Errorf("invalid CondType")
		}

		h.CondType = _type
		return nil
	}
}

func WithIndex(index *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = index
		return nil
	}
}

func WithArgumentMap(argMap *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if argMap == nil {
			if must {
				return fmt.Errorf("invalid argmap")
			}
			return nil
		}
		var r interface{}
		if err := json.Unmarshal([]byte(*argMap), &r); err != nil {
			return err
		}

		h.ArgumentMap = argMap
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		h.Conds = &crud.Conds{}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.ID.Op, Val: id}
		}

		if conds.TestCaseID != nil {
			id, err := uuid.Parse(conds.GetTestCaseID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TestCaseID = &cruder.Cond{Op: conds.TestCaseID.Op, Val: id}
		}

		if conds.CondTestCaseID != nil {
			id, err := uuid.Parse(conds.GetCondTestCaseID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CondTestCaseID = &cruder.Cond{Op: conds.CondTestCaseID.Op, Val: id}
		}

		if conds.CondType != nil {
			h.Conds.CondType = &cruder.Cond{Op: conds.CondType.Op, Val: conds.CondType}
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
