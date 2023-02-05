package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/auth/v1/authv1grpc"
	authv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/auth/v1"
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthV1Controller struct {
	svc        userservice.Serv
	jwtManager *auth.Manager
	authv1grpc.UnimplementedAuthServiceServer
}

func NewAuthV1Controller(svc userservice.Serv, jwtManager *auth.Manager) *AuthV1Controller {
	return &AuthV1Controller{svc: svc, jwtManager: jwtManager}
}

func (a AuthV1Controller) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	if request.GetUsername() == "" || request.GetPassword() == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Username and password should not be empty",
		)
	}

	usr, err := a.svc.GetByUsername(ctx, request.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot retrieve user: %v", err)
	}
	if usr == nil || !usr.IsCorrectPassword(request.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := a.jwtManager.Generate(ctx, usr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	return &authv1.LoginResponse{AccessToken: token}, nil
}
