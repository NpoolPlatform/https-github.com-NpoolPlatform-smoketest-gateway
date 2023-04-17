package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testplan"
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
	}
}

// Fields of the TestPlan.
func (TestPlan) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			String("state").
			Optional().
			Default(testplan.TestPlanState_DefaultTestPlanState.String()),
		field.
			UUID("owner_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("responsible_user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Uint32("failed_test_cases_count").
			Optional().
			Default(0),
		field.
			Uint32("passed_test_cases_count").
			Optional().
			Default(0),
		field.
			Uint32("skipped_test_cases_count").
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
			String("test_result").
			Optional().
			Default(testplan.TestResultState_DefaultTestResultState.String()),
	}
}

// Edges of the TestPlan.
func (TestPlan) Edges() []ent.Edge {
	return nil
}
