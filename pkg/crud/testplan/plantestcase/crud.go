package planplantestcase

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/plantestcase"
	"github.com/google/uuid"
)

type Req struct {
	ID             *uuid.UUID
	TestPlanID     *uuid.UUID
	TestCaseID     *uuid.UUID
	TestCaseOutput *string
	Description    *string
	RunDuration    *uint32
	TestUserID     *uuid.UUID
	Result         *npool.TestCaseResult
	Index          *uint32
	DeletedAt      *uint32
}

func CreateSet(c *ent.PlanTestCaseCreate, req *Req) *ent.PlanTestCaseCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.TestPlanID != nil {
		c.SetTestPlanID(*req.TestPlanID)
	}
	if req.TestCaseID != nil {
		c.SetTestCaseID(*req.TestCaseID)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	if req.TestCaseOutput != nil {
		c.SetTestCaseOutput(*req.TestCaseOutput)
	}
	if req.RunDuration != nil {
		c.SetRunDuration(*req.RunDuration)
	}
	if req.TestUserID != nil {
		c.SetTestUserID(*req.TestUserID)
	}
	if req.Result != nil {
		c.SetResult(req.Result.String())
	}
	if req.Index != nil {
		c.SetIndex(*req.Index)
	}
	return c
}

func UpdateSet(u *ent.PlanTestCaseUpdateOne, req *Req) *ent.PlanTestCaseUpdateOne {
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.TestCaseOutput != nil {
		u.SetTestCaseOutput(*req.TestCaseOutput)
	}
	if req.RunDuration != nil {
		u.SetRunDuration(*req.RunDuration)
	}
	if req.Result != nil {
		u.SetResult(req.Result.String())
	}
	if req.Index != nil {
		u.SetIndex(*req.Index)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID          *cruder.Cond
	TestPlanID  *cruder.Cond
	TestPlanIDs *cruder.Cond
	TestUserID  *cruder.Cond
	Result      *cruder.Cond
}

func SetQueryConds(q *ent.PlanTestCaseQuery, conds *Conds) (*ent.PlanTestCaseQuery, error) {
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(plantestcase.ID(id))
		default:
			return nil, fmt.Errorf("invalid id field")
		}
	}

	if conds.TestPlanID != nil {
		planID, ok := conds.TestPlanID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid plan id")
		}
		switch conds.TestPlanID.Op {
		case cruder.EQ:
			q.Where(plantestcase.TestPlanID(planID))
		default:
			return nil, fmt.Errorf("invalid plan id field")
		}
	}
	if conds.TestUserID != nil {
		userID, ok := conds.TestUserID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid user id")
		}
		switch conds.TestUserID.Op {
		case cruder.EQ:
			q.Where(plantestcase.TestUserID(userID))
		default:
			return nil, fmt.Errorf("invalid user id field")
		}
	}
	if conds.Result != nil {
		result, ok := conds.Result.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid result")
		}
		switch conds.Result.Op {
		case cruder.EQ:
			q.Where(plantestcase.Result(result))
		default:
			return nil, fmt.Errorf("invalid result field")
		}
	}

	if conds.TestPlanIDs != nil {
		ids, ok := conds.TestPlanIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid testplan ids")
		}
		switch conds.TestPlanIDs.Op {
		case cruder.IN:
			q.Where(plantestcase.TestPlanIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid testplan id filed")
		}
	}
	return q, nil
}
