package testplan

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/testplan"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	Name        *string
	State       *npool.TestPlanState
	CreatedBy   *uuid.UUID
	Executor    *uuid.UUID
	Fails       *uint32
	Passes      *uint32
	Skips       *uint32
	RunDuration *uint32
	Result      *npool.TestResultState
	Deadline    *uint32
	DeletedAt   *uint32
}

func CreateSet(c *ent.TestPlanCreate, req *Req) *ent.TestPlanCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.State != nil {
		c.SetState(req.State.String())
	}
	if req.CreatedBy != nil {
		c.SetCreatedBy(*req.CreatedBy)
	}
	if req.Executor != nil {
		c.SetExecutor(*req.Executor)
	}
	if req.Fails != nil {
		c.SetFails(*req.Fails)
	}
	if req.Passes != nil {
		c.SetPasses(*req.Passes)
	}
	if req.Skips != nil {
		c.SetSkips(*req.Skips)
	}
	if req.RunDuration != nil {
		c.SetRunDuration(*req.RunDuration)
	}
	if req.Result != nil {
		c.SetResult(req.Result.String())
	}
	if req.Deadline != nil {
		c.SetDeadline(*req.Deadline)
	}
	return c
}

func UpdateSet(u *ent.TestPlanUpdateOne, req *Req) *ent.TestPlanUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Executor != nil {
		u.SetExecutor(*req.Executor)
	}
	if req.State != nil {
		u.SetState(req.State.String())
	}
	if req.Deadline != nil {
		u.SetDeadline(*req.Deadline)
	}
	if req.Passes != nil {
		u.SetPasses(*req.Passes)
	}
	if req.Fails != nil {
		u.SetFails(*req.Fails)
	}
	if req.Passes != nil {
		u.SetPasses(*req.Passes)
	}
	if req.RunDuration != nil {
		u.SetRunDuration(*req.RunDuration)
	}
	if req.Result != nil {
		u.SetResult(req.Result.String())
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID        *cruder.Cond
	CreatedBy *cruder.Cond
	Executor  *cruder.Cond
	State     *cruder.Cond
}

func SetQueryConds(q *ent.TestPlanQuery, conds *Conds) (*ent.TestPlanQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(testplan.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}
	if conds.CreatedBy != nil {
		createdBy, ok := conds.CreatedBy.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid created by")
		}
		switch conds.CreatedBy.Op {
		case cruder.EQ:
			q.Where(testplan.CreatedBy(createdBy))
		default:
			return nil, fmt.Errorf("invalid created by field")
		}
	}
	if conds.Executor != nil {
		executor, ok := conds.Executor.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid executor")
		}
		switch conds.Executor.Op {
		case cruder.EQ:
			q.Where(testplan.Executor(executor))
		default:
			return nil, fmt.Errorf("invalid executor field")
		}
	}

	if conds.State != nil {
		state, ok := conds.State.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid state")
		}
		switch conds.State.Op {
		case cruder.EQ:
			q.Where(testplan.State(state))
		default:
			return nil, fmt.Errorf("invalid deprecated field")
		}
	}
	return q, nil
}
