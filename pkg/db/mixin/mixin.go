package mixin

import (
	"entgo.io/ent"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/ent/privacy"
	"github.com/NpoolPlatform/smoketest-middleware/pkg/db/rule"
)

func (TimeMixin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (TimeMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTimeRule(),
		},
	}
}
