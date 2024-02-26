package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/scoretrak/scoretrak/internal/entities/mixins"
)

// Property holds the schema definition for the Property entity.
type Property struct {
	ent.Schema
}

// Fields of the Property.
func (Property) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Immutable(),
		//field.Enum("key").Values("ip", "host"),
		field.String("value"),
		field.String("host_service_id").Immutable(),
	}
}

// Edges of the Property.
func (Property) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("hostservice", HostService.Type).Field("host_service_id").Ref("properties").Unique().Required().Immutable(),
	}
}

func (Property) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.TimeMixin{},
	}
}
