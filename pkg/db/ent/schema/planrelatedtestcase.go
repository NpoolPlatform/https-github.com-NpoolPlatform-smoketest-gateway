package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	planrelatedtestcase "github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// PlanRelatedTestCase holds the schema definition for the PlanRelatedTestCase entity.
type PlanRelatedTestCase struct {
	ent.Schema
}

func (PlanRelatedTestCase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the PlanRelatedTestCase.
func (PlanRelatedTestCase) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
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
			String("test_case_output").
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
			String("test_case_result").
			Optional().
			Default(planrelatedtestcase.TestCaseResult_DefaultTestCaseResult.String()),
	}
}

// Edges of the PlanRelatedTestCase.
func (PlanRelatedTestCase) Edges() []ent.Edge {
	return nil
}
