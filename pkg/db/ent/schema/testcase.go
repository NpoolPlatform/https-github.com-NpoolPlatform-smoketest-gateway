package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// TestCase holds the schema definition for the TestCase entity.
type TestCase struct {
	ent.Schema
}

func (TestCase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the TestCase.
func (TestCase) Fields() []ent.Field {
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
			String("description").
			Optional().
			Default(""),
		field.
			UUID("module_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("api_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("arguments").
			Optional().
			Default(""),
		field.
			Text("arg_type_description").
			Optional().
			Default(""),
		field.
			String("expectation_result").
			Optional().
			Default(""),
		field.
			String("test_case_type").
			Optional().
			Default(testcase.TestCaseType_DefaultTestCaseType.String()),
		field.
			Bool("deprecated").
			Optional().
			Default(false),
	}
}

// Edges of the TestCase.
func (TestCase) Edges() []ent.Edge {
	return nil
}
