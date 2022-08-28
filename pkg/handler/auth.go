package handler

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	authv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/auth/v1"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthController struct {
	svc        userservice.Serv
	jwtManager *auth.Manager
	authv1.UnimplementedAuthServiceServer
}

func NewAuthController(svc userservice.Serv, jwtManager *auth.Manager) *AuthController {
	return &AuthController{svc: svc, jwtManager: jwtManager}
}

func (a AuthController) Login(ctx context.Context, request *authv1.LoginRequest) (*authv1.LoginResponse, error) {
	ctx, span := otel.Tracer("scoretrak/master").Start(ctx, "Login", trace.WithAttributes(attribute.String("username", request.Username)))
	defer span.End()
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
	return &authv1.LoginResponse{AccessToken: token}, nil
}
