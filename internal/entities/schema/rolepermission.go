package schema

//// RolePermission holds the schema definition for the RolePermission entity.
//type RolePermission struct {
//	ent.Schema
//}
//
//// Fields of the RolePermission.
//func (RolePermission) Fields() []ent.Field {
//	return []ent.Field{
//		field.Enum("permission").GoType(permission.Permission("")),
//	}
//}
//
//// Edges of the RolePermission.
//func (RolePermission) Edges() []ent.Edge {
//	return []ent.Edge{
//		//edge.From("role", Role.Type).Ref("permissions").Unique().Required().Immutable(),
//		edge.From("role", Role.Type).Ref("permissions").Unique().Required(),
//	}
//}
//
//// Mixin of the RolePermission.
//func (RolePermission) Mixin() []ent.Mixin {
//	return []ent.Mixin{
//		UlidMixin{},
//		TimeMixin{},
//	}
//}
