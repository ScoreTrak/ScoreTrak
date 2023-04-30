package schema

//
//import (
//	"entgo.io/ent"
//	"entgo.io/ent/schema/edge"
//	"entgo.io/ent/schema/field"
//)
//
//// CompetitionMixin holds the schema definition for the CompetitionMixin.
//type CompetitionMixin struct {
//	ent.Schema
//}
//
//// Fields of the CompetitionMixin.
//func (CompetitionMixin) Fields() []ent.Field {
//	return []ent.Field{
//		field.String("competition_id").Immutable(),
//	}
//}
//
//// Edges of the CompetitionMixin.
//func (CompetitionMixin) Edges() []ent.Edge {
//	return []ent.Edge{
//		edge.From("competition", Competition.Type).Field("competition_id").Unique().Required().Immutable(),
//	}
//}
//
//// Mixins of the CompetitionMixin.
//func (CompetitionMixin) Mixin() []ent.Mixin {
//	return []ent.Mixin{}
//}
//
//// Policy of the CompetitionMixin
//func (CompetitionMixin) Policy() {
//
//}
