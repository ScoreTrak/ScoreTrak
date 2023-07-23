package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
)

// Check holds the schema definition for the Check entity.
type Check struct {
	ent.Schema
}

// Fields of the Check.
func (Check) Fields() []ent.Field {
	return []ent.Field{
		field.String("log").Optional(),
		field.String("error").Optional(),
		field.Bool("passed"),
		field.String("round_id").Immutable(),
		field.String("host_service_id").Immutable(),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("round", Round.Type).Field("round_id").Ref("checks").Unique().Required().Immutable(),
		edge.From("hostservice", HostService.Type).Field("host_service_id").Ref("checks").Unique().Required().Immutable(),
	}
}

// Mixin of the Check.
func (Check) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.TimeMixin{},
	}
}
