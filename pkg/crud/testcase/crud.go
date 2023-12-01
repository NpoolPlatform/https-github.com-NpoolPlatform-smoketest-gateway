package testcase

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testcase"
	"github.com/google/uuid"
)

type Req struct {
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
	DeletedAt     *uint32
}

func CreateSet(c *ent.TestCaseCreate, req *Req) *ent.TestCaseCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.ModuleID != nil {
		c.SetModuleID(*req.ModuleID)
	}
	if req.ApiID != nil {
		c.SetAPIID(*req.ApiID)
	}
	if req.Input != nil {
		c.SetInput(*req.Input)
	}
	if req.InputDesc != nil {
		c.SetInputDesc(*req.InputDesc)
	}
	if req.Expectation != nil {
		c.SetExpectation(*req.Expectation)
	}
	if req.OutputDesc != nil {
		c.SetOutputDesc(*req.OutputDesc)
	}
	if req.TestCaseType != nil {
		c.SetTestCaseType(req.TestCaseType.String())
	}
	if req.TestCaseClass != nil {
		c.SetTestCaseClass(req.TestCaseClass.String())
	}
	if req.Deprecated != nil {
		c.SetDeprecated(*req.Deprecated)
	}
	return c
}

func UpdateSet(u *ent.TestCaseUpdateOne, req *Req) *ent.TestCaseUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.Input != nil {
		u.SetInput(*req.Input)
	}
	if req.InputDesc != nil {
		u.SetInputDesc(*req.InputDesc)
	}
	if req.Expectation != nil {
		u.SetExpectation(*req.Expectation)
	}
	if req.OutputDesc != nil {
		u.SetOutputDesc(*req.OutputDesc)
	}
	if req.TestCaseType != nil {
		u.SetTestCaseType(req.TestCaseType.String())
	}
	if req.TestCaseClass != nil {
		u.SetTestCaseClass(req.TestCaseClass.String())
	}
	if req.Deprecated != nil {
		u.SetDeprecated(*req.Deprecated)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID         *cruder.Cond
	EntID      *cruder.Cond
	ModuleID   *cruder.Cond
	ApiID      *cruder.Cond //nolint
	Deprecated *cruder.Cond
	IDs        *cruder.Cond
	EntIDs     *cruder.Cond
}

func SetQueryConds(q *ent.TestCaseQuery, conds *Conds) (*ent.TestCaseQuery, error) { //nolint
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(testcase.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(testcase.EntID(id))
		default:
			return nil, fmt.Errorf("invalid entid field")
		}
	}
	if conds.ModuleID != nil {
		moduleID, ok := conds.ModuleID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid module id")
		}
		switch conds.ModuleID.Op {
		case cruder.EQ:
			q.Where(testcase.ModuleID(moduleID))
		default:
			return nil, fmt.Errorf("invalid module id field")
		}
	}
	if conds.ApiID != nil {
		apiID, ok := conds.ApiID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid api id")
		}
		switch conds.ApiID.Op {
		case cruder.EQ:
			q.Where(testcase.APIID(apiID))
		default:
			return nil, fmt.Errorf("invalid api id field")
		}
	}
	if conds.Deprecated != nil {
		deprecated, ok := conds.Deprecated.Val.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid deprecated")
		}
		switch conds.Deprecated.Op {
		case cruder.EQ:
			q.Where(testcase.Deprecated(deprecated))
		default:
			return nil, fmt.Errorf("invalid deprecated field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(testcase.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid testcase id filed")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(testcase.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid testcase entid filed")
		}
	}
	return q, nil
}
