package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// CreateTimeMixin adds created at time field.
type CreateTimeMixin struct{ mixin.Schema }

// Fields of the create time mixin.
func (CreateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Default(time.Now).
			Immutable().Optional(),
	}
}

// UpdateTimeMixin adds updated at time field.
type UpdateTimeMixin struct{ mixin.Schema }

// Fields of the update time mixin.
func (UpdateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).Optional(),
	}
}

// TimeMixin composes create/update time mixin.
type TimeMixin struct{ mixin.Schema }

// Fields of the time mixin.
func (TimeMixin) Fields() []ent.Field {
	return append(
		CreateTimeMixin{}.Fields(),
		UpdateTimeMixin{}.Fields()...,
	)
}

//func FilterTeamRule()
