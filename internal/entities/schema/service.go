package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/scoretrak/scoretrak/internal/entities/mixins"
	"github.com/scoretrak/scoretrak/pkg/scorer"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").GoType(scorer.ServiceType("")),
		field.Int("weight").Positive().Default(1),
		field.Int("point_boost").NonNegative().Default(0),
		field.Int("round_frequency").Default(1).Positive().Comment("The number of times to score every x round. If you round freq is 1, it will be scored every round. If it is 3, it will be scored, every three rounds. Ex. 1 not scored, 2 not scored, 3 scored, 4 not scored."),
		field.Int("round_delay").Default(0).NonNegative().Comment("The number of rounds to delay a round by. Round numbers start at 1. If your round delay is 0, it will start scoring in round 1. If you round delay is 3, it will start scoring on the 4th round. "),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hostservices", HostService.Type),
		edge.To("hostservicereports", HostServiceReport.Type),
	}
}

func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.NameMixin{},
		mixins.PauseMixin{},
		mixins.HideMixin{},
		mixins.TimeMixin{},
	}
}
