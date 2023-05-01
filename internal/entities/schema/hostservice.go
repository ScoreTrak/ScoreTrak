package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// HostService holds the schema definition for the HostService entity.
type HostService struct {
	ent.Schema
}

// Fields of the HostService.
func (HostService) Fields() []ent.Field {
	return []ent.Field{
		field.Int("weight").Default(1),
		field.Int("point_boost").Default(1),
		field.Int("round_units").Default(1),
		field.Int("round_delay").Default(1),
		field.String("service_id"),
		field.String("host_id").Immutable(),
		field.String("team_id").Immutable(),
		//field.String("competition_id").Immutable(),
	}
}

// Edges of the HostService.
func (HostService) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checks", Check.Type),
		edge.To("properties", Property.Type),
		edge.From("service", Service.Type).Ref("hostservices").Field("service_id").Unique().Required(),
		edge.From("host", Host.Type).Ref("hostservices").Field("host_id").Unique().Required().Immutable(),
		edge.From("team", Team.Type).Ref("hostservices").Field("team_id").Unique().Required().Immutable(),
		//edge.From("competition", Competition.Type).Ref("hostservices").Field("competition_id").Unique().Required().Immutable(),
	}
}

// Mixins of the HostService.
func (HostService) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		NameMixin{},
		PauseMixin{},
		HideMixin{},
	}
}
