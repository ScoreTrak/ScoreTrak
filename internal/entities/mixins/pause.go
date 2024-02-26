package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type PauseMixin struct {
	mixin.Schema
}

func (PauseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("pause").Optional().Default(false),
	}
}
