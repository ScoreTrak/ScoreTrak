package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
)

// TeamReport holds the schema definition for the TeamReport entity.
type TeamReport struct {
	ent.Schema
}

// Fields of the TeamReport.
func (TeamReport) Fields() []ent.Field {
	return []ent.Field{
		field.Int("points"),
		field.String("team_id").Immutable(),
	}
}

// Edges of the TeamReport.
func (TeamReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", Team.Type).Field("team_id").Ref("teamreport").Unique().Required().Immutable(),
		edge.To("hostservicereports", HostServiceReport.Type),
	}
}

// Mixin of the TeamReport.
func (TeamReport) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.TimeMixin{},
	}
}
