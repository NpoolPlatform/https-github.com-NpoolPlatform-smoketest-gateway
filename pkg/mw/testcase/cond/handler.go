package cond

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	mgrpb "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
	constant "github.com/NpoolPlatform/smoketest-middleware/pkg/const"
	crud "github.com/NpoolPlatform/smoketest-middleware/pkg/crud/testcase/cond"
	testcasemw "github.com/NpoolPlatform/smoketest-middleware/pkg/mw/testcase"
	"github.com/google/uuid"
)

type Handler struct {
	ID             *uuid.UUID
	TestCaseID     *uuid.UUID
	CondTestCaseID *uuid.UUID
	CondType       *mgrpb.CondType
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

func WithTestCaseID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		type TestCaseHandler struct {
			testcasemw.Handler
		}

		testcase := &TestCaseHandler{}
		testcase.ID = &_id

		if exist, err := testcase.ExistTestCase(ctx); !exist {
			return err
		}

		h.TestCaseID = &_id
		return nil
	}
}

func WithCondTestCaseID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}

		type TestCaseHandler struct {
			testcasemw.Handler
		}

		testcase := &TestCaseHandler{}
		testcase.ID = &_id

		if exist, err := testcase.ExistTestCase(ctx); !exist {
			return err
		}

		h.CondTestCaseID = &_id
		return nil
	}
}

func WithCondType(_type *mgrpb.CondType) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if _type == nil {
			return nil
		}
		switch *_type {
		case mgrpb.CondType_PreCondition:
		case mgrpb.CondType_Cleaner:
		default:
			return fmt.Errorf("invalid CondType")
		}

		h.CondType = _type
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

func WithArgumentMap(argMap *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if argMap == nil {
			return nil
		}
		h.ArgumentMap = argMap
		return nil
	}
}

func WithConds(conds *mgrpb.Conds) func(context.Context, *Handler) error {
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
			h.Conds.ID = &cruder.Cond{Op: h.Conds.ID.Op, Val: id}
		}

		if conds.TestCaseID != nil {
			id, err := uuid.Parse(conds.GetTestCaseID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.TestCaseID = &cruder.Cond{Op: h.Conds.TestCaseID.Op, Val: id}
		}

		if conds.CondTestCaseID != nil {
			id, err := uuid.Parse(conds.GetCondTestCaseID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CondTestCaseID = &cruder.Cond{Op: h.Conds.CondTestCaseID.Op, Val: id}
		}

		if conds.CondType != nil {
			h.Conds.CondType = &cruder.Cond{Op: h.Conds.CondType.Op, Val: conds.CondType}
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
