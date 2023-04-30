package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Report holds the schema definition for the Report entity.
type Report struct {
	ent.Schema
}

// Fields of the Report.
func (Report) Fields() []ent.Field {
	return []ent.Field{
		field.String("log"),
		field.String("error"),
		field.String("competition_id").Immutable(),
	}
}

// Edges of the Report.
func (Report) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("competition", Competition.Type).Ref("reports").Field("competition_id").Unique().Required().Immutable(),
	}
}

func (Report) Mixins() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
