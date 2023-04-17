package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Check holds the schema definition for the Check entity.
type Check struct {
	ent.Schema
}

// Fields of the Check.
func (Check) Fields() []ent.Field {
	return []ent.Field{
		field.String("log"),
		field.String("error"),
		field.Bool("passed"),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("rounds", Round.Type).Ref("checks").Unique().Required(),
		edge.From("services", Service.Type).Ref("checks").Unique().Required(),
	}
}

// Mixins of the Check.
func (Check) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		PauseMixin{},
		HideMixin{},
		CompetitonMixin{},
		//TeamMixin{},
	}
}
