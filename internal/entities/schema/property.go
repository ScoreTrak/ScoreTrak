package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
		field.String("host_service_id"),
		field.String("team_id").Immutable(),
		//field.String("competition_id").Immutable(),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hostservice", HostService.Type).Ref("properties").Field("host_service_id").Unique().Required(),
		edge.From("team", Team.Type).Ref("properties").Field("team_id").Unique().Required().Immutable(),
		//edge.From("competition", Competition.Type).Ref("properties").Field("competition_id").Unique().Required().Immutable(),
	}
}

func (Property) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
