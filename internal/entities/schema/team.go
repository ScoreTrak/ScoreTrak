package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").NonNegative(),
		field.String("competition_id").Immutable(),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hosts", Host.Type),
		edge.To("hostservices", HostService.Type),
		edge.To("checks", Check.Type),
		edge.To("properties", Property.Type),
		edge.From("competition", Competition.Type).Ref("teams").Field("competition_id").Unique().Required().Immutable(),
	}
}

// Mixins of the Team.
func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		NameMixin{},
		PauseMixin{},
		HideMixin{},
	}
}
