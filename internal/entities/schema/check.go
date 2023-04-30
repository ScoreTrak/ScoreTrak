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
		field.String("round_id").Immutable(),
		field.String("host_service_id").Immutable(),
		field.String("team_id").Immutable(),
		//field.String("competition_id").Immutable(),
	}
}

// Edges of the Check.
func (Check) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rounds", Round.Type).Field("round_id").Unique().Required().Immutable(),
		edge.From("hostservice", HostService.Type).Ref("checks").Field("host_service_id").Unique().Required().Immutable(),
		edge.From("team", Team.Type).Ref("checks").Field("team_id").Unique().Required().Immutable(),
		//edge.From("competition", Competition.Type).Ref("checks").Field("competition_id").Unique().Required().Immutable(),
	}
}

// Mixins of the Check.
func (Check) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		PauseMixin{},
		HideMixin{},
	}
}
