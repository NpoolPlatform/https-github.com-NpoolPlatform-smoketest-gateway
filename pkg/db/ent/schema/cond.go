package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// Cond holds the schema definition for the Cond entity.
type Cond struct {
	ent.Schema
}

func (Cond) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Cond.
func (Cond) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("cond_type").
			Optional().
			Default(cond.CondType_DefaultCondType.String()),
		field.
			UUID("test_case_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("cond_test_case_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Text("argument_map").
			Optional().
			Default(""),
		field.
			Uint32("index").
			Optional().
			Default(0),
	}
}

// Edges of the Cond.
func (Cond) Edges() []ent.Edge {
	return nil
}
