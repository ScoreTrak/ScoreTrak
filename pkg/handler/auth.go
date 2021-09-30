package handler

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/user_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	svc        user_service.Serv
	jwtManager *auth.Manager
	authpb.UnimplementedAuthServiceServer
}

func NewAuthController(svc user_service.Serv, jwtManager *auth.Manager) *AuthController {
	return &AuthController{svc: svc, jwtManager: jwtManager}
}

func (a AuthController) Login(ctx context.Context, request *authpb.LoginRequest) (*authpb.LoginResponse, error) {
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
	token, err := a.jwtManager.Generate(usr)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}
	return &authpb.LoginResponse{AccessToken: token}, nil
}
