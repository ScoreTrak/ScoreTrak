package rule

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/privacy"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/session"
)

func DenyIfNoSession() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		_, ok := session.FromContext(ctx)
		if !ok {
			return privacy.Denyf("session is missing")
		}

		return privacy.Skip
	})
}

//func FilterCompetitionRule() privacy.QueryRule {
//	type CompetitionsFilter interface {
//		WhereCompetitionID(p entql.StringP)
//	}
//	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
//		s, ok := session.FromContext(ctx)
//
//	})
//}
