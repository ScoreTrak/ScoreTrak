package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/scoretrak/scoretrak/internal/entities/mixins"
	"time"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.Int("round_number").Unique().Positive(),
		field.Int("num_of_intended_checks").Optional(),
		field.Enum("status").Values("started", "ongoing", "calculated", "finished", "incomplete").Default("started"),
		field.Time("started_at").Default(time.Now()),
		field.Time("finished_at").Optional(),
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checks", Check.Type),
	}
}

func (Round) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.TimeMixin{},
	}
}
