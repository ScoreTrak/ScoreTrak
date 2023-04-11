package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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
		mixin.Time{},
		CompetitonMixin{},
	}
}
