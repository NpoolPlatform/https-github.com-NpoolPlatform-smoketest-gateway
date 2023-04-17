package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/smoketest/mgr/v1/testcase/cond"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// RelatedTestCase holds the schema definition for the RelatedTestCase entity.
type RelatedTestCase struct {
	ent.Schema
}

func (RelatedTestCase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the RelatedTestCase.
func (RelatedTestCase) Fields() []ent.Field {
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
			UUID("related_test_case_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			String("arguments_transfer").
			Optional().
			Default(""),
		field.
			Uint32("index").
			Optional().
			Default(0),
	}
}

// Edges of the RelatedTestCase.
func (RelatedTestCase) Edges() []ent.Edge {
	return nil
}
