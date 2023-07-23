package schema

import (
	"entgo.io/contrib/entoas"
	"entgo.io/ent/schema"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities/mixins"
	"regexp"

	"entgo.io/ent"
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
		field.Int("round_duration").Optional().Default(60),
		field.String("current_round_id").Optional().Nillable().Comment("Most recently completed round"),
		field.Bool("viewable_to_public").Nillable().Optional(),
		field.Bool("ignore_incomplete_round_in_scoring").Optional(),
		field.Time("to_be_started_at").Nillable().Optional(),
		field.Time("started_at").Nillable().Optional(),
		field.Time("finished_at").Nillable().Optional(),
	}
}

// Edges of the Competition.
func (Competition) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the Competition.
func (Competition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		//HideMixin{},
		mixins.PauseMixin{},
		mixins.SingleMixin{},
		mixins.TimeMixin{},
	}
}

func (Competition) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entoas.ListOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
		entoas.DeleteOperation(entoas.OperationPolicy(entoas.PolicyExclude)),
	}
}

//func (Competition) Policy() ent.Policy {
//	return privacy.Policy{
//		Query: privacy.QueryPolicy{
//			rule.DenyIfNoUser(),
//			privacy.AlwaysAllowRule(),
//		},
//		Mutation: privacy.MutationPolicy{
//			rule.AllowIfAdmin(),
//			privacy.AlwaysDenyRule(),
//		},
//	}
//}
