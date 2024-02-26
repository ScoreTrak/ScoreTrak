package rule

//import (
//	"context"
//	"entgo.io/ent/entql"
//	"github.com/scoretrak/scoretrak/internal/entities/privacy"
//)
//
//func FilterHiddenRule() privacy.QueryMutationRule {
//	type HiddenFilter interface {
//		WhereHidden(p entql.BoolP)
//	}
//	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
//		hf, ok := f.(HiddenFilter)
//		if !ok {
//			return privacy.Denyf("unexpected filter type")
//		}
//		hf.WhereHidden(entql.BoolEQ(false))
//		return privacy.Skip
//	})
//}
