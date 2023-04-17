package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("index").Optional().NonNegative(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
		edge.To("hosts", Host.Type),
	}
}

// Mixins of the Team.
func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		PauseMixin{},
		HideMixin{},
		CompetitonMixin{},
	}
}

type TeamMixin struct {
	mixin.Schema
}

func (TeamMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("team_id"),
	}
}

func (TeamMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("team", Team.Type).Field("team_id").Unique().Required(),
	}
}
