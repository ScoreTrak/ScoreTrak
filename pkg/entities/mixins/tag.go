package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type TagMixin struct {
	mixin.Schema
}

func (TagMixin) Fields() []ent.Field {
	return []ent.Field{
		//field.String("tags").Optional().Nillable().Match(regexp.MustCompile("(,?\\w+=\\w+)*")),
		field.String("tags").Optional().Nillable(),
	}
}
