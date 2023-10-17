package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/smoketest/mw/v1/testcase"
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
			Text("input").
			Optional().
			Default(""),
		field.
			Text("input_desc").
			Optional().
			Default(""),
		field.
			Text("expectation").
			Optional().
			Default(""),
		field.
			Text("output_desc").
			Optional().
			Default(""),
		field.
			String("test_case_type").
			Optional().
			Default(testcase.TestCaseType_DefaultTestCaseType.String()),
		field.
			String("test_case_class").
			Optional().
			Default(testcase.TestCaseClass_DefaultTestCaseClass.String()),
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
