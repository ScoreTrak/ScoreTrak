package server

import (
	api_stub "github.com/scoretrak/scoretrak/internal/api-stub"
	"github.com/scoretrak/scoretrak/pkg/server/handler"
)

func NewApiServer(h *handler.Handler, sh *ApiTokenSecurityHandler) (*api_stub.Server, error) {
	s, err := api_stub.NewServer(h, sh, api_stub.WithPathPrefix("/api"))
	if err != nil {
		return nil, err
	}

	return s, nil
}
