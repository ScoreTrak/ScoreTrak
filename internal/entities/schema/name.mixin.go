package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"regexp"
)

type NameMixin struct {
	mixin.Schema
}

func (NameMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Match(regexp.MustCompile("^[a-z0-9_]*$")).MaxLen(32),
		field.String("display_name").Match(regexp.MustCompile("^[a-zA-Z0-9\\s]*$")).MaxLen(64),
	}
}
