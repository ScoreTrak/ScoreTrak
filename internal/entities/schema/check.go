package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/scoretrak/scoretrak/internal/entities/mixins"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
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
		field.Enum("outcome_status").Values(string(outcome.OUTCOME_STATUS_PASSED), string(outcome.OUTCOME_STATUS_FAILED), string(outcome.OUTCOME_STATUS_INVALID)).Optional(),
		field.Enum("progress_status").Values("started", "finished").Default("started"),
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
