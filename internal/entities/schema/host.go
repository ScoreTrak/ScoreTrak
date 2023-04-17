package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Host holds the schema definition for the Host entity.
type Host struct {
	ent.Schema
}

// Fields of the Host.
func (Host) Fields() []ent.Field {
	return []ent.Field{
		field.String("address").Match(regexp.MustCompile("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$")),
		field.String("address_list_range"),
		field.Bool("editable"),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type),
		edge.From("host_group", HostGroup.Type).Ref("hosts").Unique().Required(),
	}
}

func (Host) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		PauseMixin{},
		HideMixin{},
		CompetitonMixin{},
		TeamMixin{},
	}
}
