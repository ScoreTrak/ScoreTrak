package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Competition holds the schema definition for the Competition entity.
type Competition struct {
	ent.Schema
}

// Fields of the Competition.
func (Competition) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Match(regexp.MustCompile("^[a-z0-9_]*$")).MinLen(4).MaxLen(32).Unique(),
		field.String("display_name").Match(regexp.MustCompile("^[a-zA-Z0-9\\s]*$")).MinLen(4).MaxLen(64),
		//field.Int("round_duration").Optional(),
		field.Bool("viewable_to_public").Nillable().Optional(),
		field.Time("to_be_started_at").Nillable().Optional(),
		field.Time("started_at").Nillable().Optional(),
		field.Time("finished_at").Nillable().Optional(),
	}
}

// Edges of the Competition.
func (Competition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teams", Team.Type),
		edge.To("services", Service.Type),
		edge.To("reports", Report.Type),
		edge.To("rounds", Round.Type),
		//edge.To("checks", Check.Type),
		//edge.To("hosts", Host.Type),
		//edge.To("hostservices", HostService.Type),
		//edge.To("properties", Property.Type),
	}
}

// Mixins of the Competition.
func (Competition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		HideMixin{},
		PauseMixin{},
	}
}
