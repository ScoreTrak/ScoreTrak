package schema

//// Role holds the schema definition for the Role entity.
//type Role struct {
//	ent.Schema
//}
//
//// Fields of the Role.
//func (Role) Fields() []ent.Field {
//	return []ent.Field{
//		field.String("name"),
//	}
//}
//
//// Edges of the Role.
//func (Role) Edges() []ent.Edge {
//	return []ent.Edge{
//		edge.To("permissions", RolePermission.Type),
//	}
//}
//
//// Mixin of the Role.
//func (Role) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		UlidMixin{},
//		TimeMixin{},
//	}
//}
