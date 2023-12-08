package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/mixin"
)

// Module holds the schema definition for the Module entity.
type Module struct {
	ent.Schema
}

func (Module) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Module.
func (Module) Fields() []ent.Field {
	return []ent.Field{
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
