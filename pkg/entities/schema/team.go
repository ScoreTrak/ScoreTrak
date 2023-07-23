package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").NonNegative().Unique(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hosts", Host.Type),
		edge.To("hostservices", HostService.Type),
		edge.To("teamreport", TeamReport.Type).Unique(),
		edge.To("hostservicereports", HostServiceReport.Type),
	}
}

// Mixin of the Team.
func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.NameMixin{},
		mixins.PauseMixin{},
		mixins.HideMixin{},
		mixins.TimeMixin{},
	}
}
