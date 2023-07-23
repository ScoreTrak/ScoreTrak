package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
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
		field.String("team_id"),
	}
}

// Edges of the Host.
func (Host) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hostservices", HostService.Type),
		edge.From("team", Team.Type).Field("team_id").Ref("hosts").Unique().Required(),
	}
}

func (Host) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		//NameMixin{},
		mixins.PauseMixin{},
		mixins.HideMixin{},
		mixins.TimeMixin{},
	}
}
