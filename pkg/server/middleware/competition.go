package middleware

import (
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"net/http"
)

type CompetitionMiddleware struct {
	dbClient *entities.Client
}

func NewCompetitionMiddleware(dbClient *entities.Client) *CompetitionMiddleware {
	return &CompetitionMiddleware{dbClient: dbClient}
}

func (c *CompetitionMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		//s, ok := session.FromContext(context.Background())
		//if !ok {
		//	return
		//}

	})
}
