package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// Module holds the schema definition for the Module entity.
type Module struct {
	ent.Schema
}

func (Module) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Module.
func (Module) Fields() []ent.Field {
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
	}
}

// Edges of the Module.
func (Module) Edges() []ent.Edge {
	return nil
}
