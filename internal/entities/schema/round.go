package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Round holds the schema definition for the Round entity.
type Round struct {
	ent.Schema
}

// Fields of the Round.
func (Round) Fields() []ent.Field {
	return []ent.Field{
		field.Int("round_number"),
		field.String("note"),
		field.String("err"),
		field.Time("started_at"),
		field.Time("finished_at"),
		field.String("competition_id").Immutable(),
	}
}

// Edges of the Round.
func (Round) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("checks", Check.Type),
		edge.From("competition", Competition.Type).Ref("rounds").Field("competition_id").Unique().Required().Immutable(),
	}
}

func (Round) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}