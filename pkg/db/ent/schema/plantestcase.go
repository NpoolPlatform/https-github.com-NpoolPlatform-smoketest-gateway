package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	plantestcase "github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testplan/plantestcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// PlanTestCase holds the schema definition for the PlanTestCase entity.
type PlanTestCase struct {
	ent.Schema
}

func (PlanTestCase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the PlanTestCase.
func (PlanTestCase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("test_plan_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("test_case_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Text("input").
			Optional().
			Default(""),
		field.
			Text("output").
			Optional().
			Default(""),
		field.
			String("description").
			Optional().
			Default(""),
		field.
			UUID("test_user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("run_duration").
			Optional().
			Default(0),
		field.
			String("result").
			Optional().
			Default(plantestcase.TestCaseResult_DefaultTestCaseResult.String()),
		field.
			Uint32("index").
			Optional().
			Default(0),
	}
}

// Edges of the PlanTestCase.
func (PlanTestCase) Edges() []ent.Edge {
	return nil
}
