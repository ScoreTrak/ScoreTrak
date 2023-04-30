package schema

import (
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
		//field.String("address").Match(regexp.MustCompile("^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$")),
		field.String("address"),
		//field.Bool("editable").Default(false).Optional(),
		field.String("team_id").Immutable(),
		//field.String("competition_id").Immutable(),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hostservices", HostService.Type),
		edge.From("team", Team.Type).Ref("hosts").Field("team_id").Unique().Required().Immutable(),
		//edge.From("competition", Competition.Type).Ref("hosts").Field("competition_id").Unique().Required().Immutable(),
	}
}

func (Host) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		PauseMixin{},
		HideMixin{},
	}
}
