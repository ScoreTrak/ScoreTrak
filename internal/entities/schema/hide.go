package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type HideMixin struct {
	mixin.Schema
}

func (HideMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("hidden"),
	}
}
