package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("display_name"),
		field.Int("weight"),
		field.Int("point_boost"),
		field.Int("round_units"),
		field.Int("round_delay"),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hosts", Host.Type).Ref("services").Unique(),
		edge.To("checks", Check.Type),
		edge.To("properties", Property.Type),
	}
}

// Mixins of the Service.
func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		PauseMixin{},
		HideMixin{},
		CompetitonMixin{},
		TeamMixin{},
	}
}
