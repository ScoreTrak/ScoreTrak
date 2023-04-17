package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	//"github.com/ScoreTrak/ScoreTrak/internal/entities/privacy"
	"github.com/gofrs/uuid"
	"regexp"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").Match(regexp.MustCompile("/^[a-z0-9_]+$/")),
		field.UUID("ory_id", uuid.UUID{}).StorageKey("oid").Immutable().Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teams", Team.Type).Ref("users"),
		edge.From("competitions", Competition.Type).Ref("users"),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TimeMixin{},
	}
}

//func (User) Policy() ent.Policy {
//	return privacy.Policy{
//		Mutation: privacy.MutationPolicy{
//			privacy.AlwaysDenyRule(),
//		},
//		Query: privacy.QueryPolicy{
//			privacy.AlwaysAllowRule(),
//		},
//	}
//}
