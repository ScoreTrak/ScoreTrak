package connect

import "net/http"

type Route struct {
	Path    string
	Handler http.Handler
}

func NewRoute(path string, handler http.Handler) Route {
	return Route{
		Path:    path,
		Handler: handler,
	}
}
