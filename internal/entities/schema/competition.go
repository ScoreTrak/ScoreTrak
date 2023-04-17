package schema

import (
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Competition holds the schema definition for the Competition entity.
type Competition struct {
	ent.Schema
}

// Fields of the Competition.
func (Competition) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Match(regexp.MustCompile("^[a-z0-9_]{4,32}$")).Unique(),
		field.String("display_name").Unique(),
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
		edge.To("users", User.Type),
	}
}

//func (Competition) Policy() ent.Policy {
//  return privacy.Policy{
//    Query: privacy.QueryPolicy{
//      //rule.DenyIfNoViewer(),
//      //rule.AllowIfAdmin(),
//      rule.DenyIfNoSession(),
//    },
//    Mutation: privacy.MutationPolicy{
//      //rule.AllowIfAdmin(),
//      //privacy.AlwaysDenyRule(),
//    },
//  }
//}

// Mixins of the Competition.
func (Competition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		BaseMixin{},
		HideMixin{},
		PauseMixin{},
	}
}

type CompetitonMixin struct {
	mixin.Schema
}

func (CompetitonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("competition_id").Immutable(),
	}
}

func (CompetitonMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("competition", Competition.Type).Field("competition_id").Unique().Required().Immutable(),
	}
}
