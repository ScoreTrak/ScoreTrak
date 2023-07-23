package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
)

// HostServiceReport holds the schema definition for the HostServiceReport entity.
type HostServiceReport struct {
	ent.Schema
}

// Fields of the HostServiceReport.
func (HostServiceReport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("points"),
		field.Bool("passing"),
		field.Time("latest_check_time"),
		field.String("host_service_id").Immutable(),
		field.String("service_id").Immutable(),
		field.String("team_id").Immutable(),
		field.String("team_report_id").Optional().Immutable(),
	}
}

// Edges of the HostServiceReport.
func (HostServiceReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hostservice", HostService.Type).Field("host_service_id").Ref("hostservicereport").Unique().Required().Immutable(),
		edge.From("service", Service.Type).Field("service_id").Ref("hostservicereports").Unique().Required().Immutable(),
		edge.From("team", Team.Type).Field("team_id").Ref("hostservicereports").Unique().Required().Immutable(),
		edge.From("teamreport", TeamReport.Type).Field("team_report_id").Ref("hostservicereports").Unique().Immutable(),
	}
}

// Mixin of the HostServiceReport.
func (HostServiceReport) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.TimeMixin{},
	}
}
