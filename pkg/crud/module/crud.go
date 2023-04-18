package module

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uuid.UUID
	Name        *string
	Description *string
}

func CreateSet(c *ent.ModuleCreate, req *Req) *ent.ModuleCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	return c
}

func UpdateSet(ctx context.Context, u *ent.ModuleUpdateOne, req *Req) (*ent.ModuleUpdateOne, error) {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	return u, nil
}

type Conds struct {
	ID   *cruder.Cond
	Name *cruder.Cond
	IDs  *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.ModuleQuery, conds *Conds) (*ent.ModuleQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(module.ID(id))
		default:
			return nil, fmt.Errorf("invalid module id field")
		}
	}
	if conds.Name != nil {
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(module.Name(conds.Name.Op))
		default:
			return nil, fmt.Errorf("invalid module name field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(module.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid module ids field")
		}
	}
	return q, nil
}