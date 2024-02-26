package mixins

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/oklog/ulid/v2"
)

// UlidMixin to be shared will all different schemas.
type UlidMixin struct {
	mixin.Schema
}

// Fields of the Mixin.
func (UlidMixin) Fields() []ent.Field {
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
