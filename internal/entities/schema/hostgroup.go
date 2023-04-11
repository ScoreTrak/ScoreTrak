package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// HostGroup holds the schema definition for the HostGroup entity.
type HostGroup struct {
	ent.Schema
}

// Fields of the HostGroup.
func (HostGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the HostGroup.
func (HostGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hosts", Host.Type),
	}
}

func (HostGroup) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		PauseMixin{},
		HideMixin{},
		CompetitonMixin{},
		TeamMixin{},
	}
}
