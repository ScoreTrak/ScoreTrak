package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

// BaseMixin to be shared will all different schemas.
type BaseMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").DefaultFunc(func() string {
			u := ulid.Make()
			return u.String()
		}).Validate(func(s string) error {
			_, err := ulid.Parse(s)
			return err
		}),
	}
}

func (BaseMixin) Mixins() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}
