package handler

//
//import (
//	"buf.build/gen/go/scoretrak/scoretrakapis/bufbuild/connect-go/scoretrak/auth/v2/authv2connect"
//	authv2 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/auth/v2"
//	"context"
//	"errors"
//	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
//	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
//	"github.com/bufbuild/connect-go"
//)
//
//type AuthV2ConnectServer struct {
//	svc        userservice.Serv
//	jwtManager *auth.Manager
//	authv2connect.UnimplementedAuthServiceHandler
//}
//
//func NewAuthV2ConnectServer(svc userservice.Serv, jwtManager *auth.Manager) *AuthV2ConnectServer {
//	return &AuthV2ConnectServer{svc: svc, jwtManager: jwtManager}
//}
//
//func (s AuthV2ConnectServer) Login(ctx context.Context, req *connect.Request[authv2.AuthServiceLoginRequest]) (*connect.Response[authv2.AuthServiceLoginResponse], error) {
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
//	return connect.NewResponse(&authv2.AuthServiceLoginResponse{AccessToken: token}), nil
//}
