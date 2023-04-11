package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Property holds the schema definition for the Property entity.
type Property struct {
	ent.Schema
}

// Fields of the Property.
func (Property) Fields() []ent.Field {
	return []ent.Field{
		//field.Enum("key").Values("IP", "Port", "Password"),
		field.String("key"),
		field.String("value"),
		field.Enum("status").Values("view", "edit", "hide").Default("view"),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("services", Service.Type).Ref("properties").Unique().Required(),
	}
}

func (Property) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		CompetitonMixin{},
		TeamMixin{},
	}
}
