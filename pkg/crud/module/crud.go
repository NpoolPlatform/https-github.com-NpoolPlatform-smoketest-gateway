package module

import (
	"fmt"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/module"
	"github.com/google/uuid"
)

type Req struct {
	ID          *uint32
	EntID       *uuid.UUID
	Name        *string
	Description *string
	DeletedAt   *uint32
}

func CreateSet(c *ent.ModuleCreate, req *Req) *ent.ModuleCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.Name != nil {
		c.SetName(*req.Name)
	}
	if req.Description != nil {
		c.SetDescription(*req.Description)
	}
	return c
}

func UpdateSet(u *ent.ModuleUpdateOne, req *Req) *ent.ModuleUpdateOne {
	if req.Name != nil {
		u.SetName(*req.Name)
	}
	if req.Description != nil {
		u.SetDescription(*req.Description)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	return u
}

type Conds struct {
	ID     *cruder.Cond
	EntID  *cruder.Cond
	Name   *cruder.Cond
	IDs    *cruder.Cond
	EntIDs *cruder.Cond
}

//nolint:nolintlint,gocyclo
func SetQueryConds(q *ent.ModuleQuery, conds *Conds) (*ent.ModuleQuery, error) {
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
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
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(module.EntID(id))
		default:
			return nil, fmt.Errorf("invalid module entid field")
		}
	}
	if conds.Name != nil {
		name, ok := conds.Name.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid name")
		}
		switch conds.Name.Op {
		case cruder.EQ:
			q.Where(module.Name(name))
		default:
			return nil, fmt.Errorf("invalid module name field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(module.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid module ids field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid entid")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(module.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid module entids field")
		}
	}
	return q, nil
}
