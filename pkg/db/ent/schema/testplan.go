package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// TestPlan holds the schema definition for the TestPlan entity.
type TestPlan struct {
	ent.Schema
}

func (TestPlan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the TestPlan.
func (TestPlan) Fields() []ent.Field {
	return []ent.Field{
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("state").
			Optional().
			Default(testplan.TestPlanState_WaitStart.String()),
		field.
			UUID("created_by", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("executor", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("fails").
			Optional().
			Default(0),
		field.
			Uint32("passes").
			Optional().
			Default(0),
		field.
			Uint32("skips").
			Optional().
			Default(0),
		field.
			Uint32("run_duration").
			Optional().
			Default(0),
		field.
			Uint32("deadline").
			Optional().
			Default(0),
		field.
			String("result").
			Optional().
			Default(testplan.TestResultState_DefaultTestResultState.String()),
	}
}

// Edges of the TestPlan.
func (TestPlan) Edges() []ent.Edge {
	return nil
}
