package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.String("competition_id").Immutable(),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("competition", Competition.Type).Ref("services").Field("competition_id").Unique().Required().Immutable(),
	}
}

func (Service) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		NameMixin{},
		PauseMixin{},
		HideMixin{},
	}
}
