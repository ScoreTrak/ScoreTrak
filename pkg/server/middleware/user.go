package middleware

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/internal/entities"
	"github.com/ScoreTrak/ScoreTrak/internal/entities/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth/session"
	"net/http"
)

type UserMiddleware struct {
	dbClient *entities.Client
}

func NewUserMiddleware(dbClient *entities.Client) *UserMiddleware {
	return &UserMiddleware{dbClient: dbClient}
}

func (c *UserMiddleware) Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		s, ok := session.FromContext(context.Background())
		if !ok {
			return
		}
		fmt.Println(s.Identity.Id)
		fmt.Println(s.Identity.Traits)
		fmt.Println(s.Identity.Credentials)
		// Check if user exist
		c.dbClient.User.Query().Where(user.ByOryID())
		u, err := c.dbClient.User.Get(context.Background(), s.Identity.Id)
		if err != nil {
			return
		}

	})
}
