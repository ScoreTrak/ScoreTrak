package handler

//import (
//	"buf.build/gen/go/scoretrak/scoretrakapis/bufbuild/connect-go/scoretrak/auth/v1/authv1connect"
//	"context"
//	"errors"
//	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
//	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
//	"github.com/bufbuild/connect-go"
//	authv1 "buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/auth/v1"
//)
//
//type AuthV1ConnectServer struct {
//	svc        userservice.Serv
//	jwtManager *auth.Manager
//	authv1connect.UnimplementedAuthServiceHandler
//}
//
//func NewAuthV1ConnectServer(svc userservice.Serv, jwtManager *auth.Manager) *AuthV1ConnectServer {
//	return &AuthV1ConnectServer{svc: svc, jwtManager: jwtManager}
//}
//
//func (s AuthV1ConnectServer) Login(ctx context.Context, req *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
//	UsernameAndPasswordEmptyError := errors.New("Username and password should not be empty")
//	CannotRetriveUserError := errors.New("Cannot retrive user")
//	IncorrectUsernamePasswordError := errors.New("Incorrect username/password")
//	CannotGenerateAccessTokenError := errors.New("Cannot generate access token")
//
//	if req.Msg.Username == "" || req.Msg.Password == "" {
//		return nil, connect.NewError(connect.CodeInvalidArgument, UsernameAndPasswordEmptyError)
//	}
//
//	usr, err := s.svc.GetByUsername(ctx, req.Msg.Username)
//	if err != nil {
//		return nil, connect.NewError(connect.CodeInternal, CannotRetriveUserError)
//	}
//	if usr == nil || !usr.IsCorrectPassword(req.Msg.Password) {
//		return nil, connect.NewError(connect.CodeNotFound, IncorrectUsernamePasswordError)
//	}
//	token, err := s.jwtManager.Generate(ctx, usr)
//	if err != nil {
//		return nil, connect.NewError(connect.CodeInternal, CannotGenerateAccessTokenError)
//	}
//	return connect.NewResponse(&authv1.LoginResponse{AccessToken: token}), nil
//}
