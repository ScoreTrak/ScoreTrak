package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
)

// HostService holds the schema definition for the HostService entity.
type HostService struct {
	ent.Schema
}

// Fields of the HostService.
func (HostService) Fields() []ent.Field {
	return []ent.Field{
		field.String("service_id").Immutable(),
		field.String("host_id").Immutable(),
		field.String("team_id").Immutable(),
	}
}

// Edges of the HostService.
func (HostService) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checks", Check.Type),
		edge.To("properties", Property.Type),
		edge.To("hostservicereport", HostServiceReport.Type).Unique(),
		edge.From("service", Service.Type).Field("service_id").Ref("hostservices").Unique().Required().Immutable(),
		edge.From("host", Host.Type).Field("host_id").Ref("hostservices").Unique().Required().Immutable(),
		edge.From("team", Team.Type).Field("team_id").Ref("hostservices").Unique().Required().Immutable(),
	}
}

// Mixin of the HostService.
func (HostService) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.NameMixin{},
		mixins.PauseMixin{},
		mixins.HideMixin{},
		mixins.TimeMixin{},
	}
}
