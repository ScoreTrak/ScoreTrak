package middleware

import "github.com/ogen-go/ogen/middleware"

type CorsMiddleware struct {
}

func (c *CorsMiddleware) Middleware(req middleware.Request, next middleware.Next) (middleware.Response, error) {

}
