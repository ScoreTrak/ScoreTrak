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
		// field.UUID("id", uuid.UUID{}).Default(uuid.NewV4),
		// field.String("id", ulid.ULID).Default(ulid.Make),
		field.String("name").Match(regexp.MustCompile("^[a-z0-9_]{4,32}$")).Unique(),
		field.String("display_name").Unique(),
		field.Float("round_duration"),
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
//	return privacy.Policy{
//		Query: privacy.QueryPolicy{
//			rule.DenyIfNoViewer(),
//			rule.AllowIfAdmin(),
//		},
//		Mutation: privacy.MutationPolicy{
//			rule.AllowIfAdmin(),
//			privacy.AlwaysDenyRule(),
//		},
//	}
//}

// Mixins of the Competition.
func (Competition) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		HideMixin{},
		PauseMixin{},
	}
}

type CompetitonMixin struct {
	mixin.Schema
}

func (CompetitonMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("competition_id").Immutable(),
	}
}

func (CompetitonMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("competition", Competition.Type).Field("competition_id").Unique().Required().Immutable(),
	}
}

//func FilterCompetitionRule() privacy.QueryMutationRule {
//	type CompetitionsFilter interface {
//		WhereCompetitionID(entql.IntP)
//	}
//
//	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
//		//view := viewer.FromContext(ctx)
//		//cid, ok := view.
//		return nil
//	})
//}
