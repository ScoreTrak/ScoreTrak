package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/auth/v2/authv2grpc"
	authv2 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/auth/v2"
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthV2Controller struct {
	svc        userservice.Serv
	jwtManager *auth.Manager
	authv2grpc.UnimplementedAuthServiceServer
}

func NewAuthV2Controller(svc userservice.Serv, jwtManager *auth.Manager) *AuthV2Controller {
	return &AuthV2Controller{svc: svc, jwtManager: jwtManager}
}

func (a AuthV2Controller) Login(ctx context.Context, request *authv2.AuthServiceLoginRequest) (*authv2.AuthServiceLoginResponse, error) {
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
	return &authv2.AuthServiceLoginResponse{AccessToken: token}, nil
}
