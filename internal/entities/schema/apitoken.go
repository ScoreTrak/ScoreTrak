package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/scoretrak/scoretrak/internal/entities/mixins"
	"github.com/scoretrak/scoretrak/pkg/auth"
	"time"
)

// ApiToken holds the schema definition for the ApiToken entity.
type ApiToken struct {
	ent.Schema
}

// Fields of the ApiToken.
func (ApiToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Match(auth.TOKEN_REGEX).Unique(),
		//field.String("role_id").Immutable(),
		field.Time("expired_at").Default(time.Now().Add(time.Hour * 720)),
		//field.String("user_id").Immutable(),
	}
}

// Edges of the ApiToken.
func (ApiToken) Edges() []ent.Edge {
	return []ent.Edge{}
}

// Mixin of the ApiToken.
func (ApiToken) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.UlidMixin{},
		mixins.TimeMixin{},
	}
}

//func (ApiToken) Privacy() ent.Policy {
//	return privacy.Policy{
//		Query: privacy.QueryPolicy{
//			rule.AllowIfAdmin(),
//			privacy.AlwaysDenyRule(),
//		},
//	}
//}
