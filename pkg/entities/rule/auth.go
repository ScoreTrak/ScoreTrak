package rule

//import (
//	"context"
//	"github.com/ScoreTrak/ScoreTrak/pkg/auth/user"
//	"github.com/ScoreTrak/ScoreTrak/pkg/entities/privacy"
//)
//
//func DenyIfNoUser() privacy.QueryMutationRule {
//	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
//		_, ok := user.FromContext(ctx)
//		if !ok {
//			return privacy.Denyf("Unauthenticated")
//		}
//		return privacy.Skip
//	})
//}
//
//func AllowIfAdmin() privacy.QueryMutationRule {
//	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
//		u, ok := user.FromContext(ctx)
//		if ok {
//			if user.IsAdmin(u) {
//				return privacy.Allow
//			}
//		}
//
//		return privacy.Skip
//	})
//}
