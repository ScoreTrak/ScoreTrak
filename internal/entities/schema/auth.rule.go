package schema

//func DenyIfNoUser() privacy.QueryMutationRule {
//	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
//		_, ok := user.FromContext(ctx)
//		if !ok {
//			return privacy.Denyf("cannot interact with resource")
//		}
//		return privacy.Skip
//	})
//}
//
//func AllowIfAdmin() privacy.QueryMutationRule {
//	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
//		u, ok := user.FromContext(ctx)
//		if !ok {
//			return privacy.Skip
//		}
//		if user.IsAdmin(u) {
//			return privacy.Allow
//		}
//		return privacy.Skip
//	})
//}
