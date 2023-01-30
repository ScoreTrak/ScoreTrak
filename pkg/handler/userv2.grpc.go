package handler

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	"github.com/gofrs/uuid"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	userv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v2"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserV2Controller struct {
	svc userservice.Serv
	userv2.UnimplementedUserServiceServer
	policyClient *policyclient.Client
}

func (p UserV2Controller) GetByUsername(ctx context.Context, request *userv2.UserServiceGetByUsernameRequest) (*userv2.UserServiceGetByUsernameResponse, error) {
	username := request.GetUsername()
	if username == "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Username was not specified",
		)
	}
	tm, err := p.svc.GetByUsername(ctx, username)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &userv2.UserServiceGetByUsernameResponse{User: ConvertUserToUserV2Pb(tm)}, nil
}

func (p UserV2Controller) GetByID(ctx context.Context, request *userv2.UserServiceGetByIDRequest) (*userv2.UserServiceGetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	tm, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &userv2.UserServiceGetByIDResponse{User: ConvertUserToUserV2Pb(tm)}, nil
}

func (p UserV2Controller) GetAll(ctx context.Context, _ *userv2.UserServiceGetAllRequest) (*userv2.UserServiceGetAllResponse, error) {
	tms, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	tmspb := make([]*userv2.User, 0, len(tms))
	for i := range tms {
		tmspb = append(tmspb, ConvertUserToUserV2Pb(tms[i]))
	}
	return &userv2.UserServiceGetAllResponse{Users: tmspb}, nil
}

func (p UserV2Controller) Delete(ctx context.Context, request *userv2.UserServiceDeleteRequest) (*userv2.UserServiceDeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &userv2.UserServiceDeleteResponse{}, nil
}

func (p UserV2Controller) Store(ctx context.Context, request *userv2.UserServiceStoreRequest) (*userv2.UserServiceStoreResponse, error) {
	usrspb := request.GetUsers()
	usrs := make([]*user.User, 0, len(usrspb))
	for i := range usrspb {
		usr, err := ConvertUserV2PBtoUser(false, usrspb[i])
		if err != nil {
			return nil, err
		}
		if usr.TeamID == uuid.Nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Team ID should not be nil",
			)
		}
		if usr.PasswordHash == "" {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Password should not be empty",
			)
		}

		if usr.Role == "" {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Role should not be empty",
			)
		}

		usrs = append(usrs, usr)
	}
	if err := p.svc.Store(ctx, usrs); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	ids := make([]*protov1.UUID, 0, len(usrs))
	for i := range usrs {
		ids = append(ids, &protov1.UUID{Value: usrs[i].ID.String()})
	}
	return &userv2.UserServiceStoreResponse{Ids: ids}, nil
}

func (p UserV2Controller) Update(ctx context.Context, request *userv2.UserServiceUpdateRequest) (*userv2.UserServiceUpdateResponse, error) {
	usrpb := request.GetUser()
	usr, err := ConvertUserV2PBtoUser(true, usrpb)
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)

	if claim.Role != user.Black {
		if claim.ID != usr.TeamID.String() {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+changingUser,
			)
		}

		if !p.policyClient.GetAllowChangingUsernamesAndPasswords() {
			return nil, status.Errorf(
				codes.PermissionDenied,
				noPermissionsTo+changingUser+"'s password",
			)
		}

		request.User = &userv2.User{Username: request.GetUser().Username, Password: request.User.GetPassword()}
	}

	err = p.svc.Update(ctx, usr)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &userv2.UserServiceUpdateResponse{}, nil
}

func NewUserV2Controller(svc userservice.Serv, policyClient *policyclient.Client) *UserV2Controller {
	return &UserV2Controller{svc: svc, policyClient: policyClient}
}

func ConvertUserV2PBtoUser(requireID bool, pb *userv2.User) (*user.User, error) {
	var id uuid.UUID
	var err error
	if pb.GetId() != nil {
		id, err = uuid.FromString(pb.GetId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				unableToParseID+": %v", err,
			)
		}
	} else if requireID {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}

	var teamID uuid.UUID
	if pb.GetTeamId() != nil {
		teamID, err = uuid.FromString(pb.GetTeamId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				unableToParseID+": %v", err,
			)
		}
	}

	var passwordHash []byte
	switch {
	case pb.GetPassword() != "" && pb.GetPasswordHash() != "":
		return nil, status.Errorf(
			codes.InvalidArgument,
			"You should provide either password or hash, but not both",
		)
	case pb.GetPassword() != "":
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(pb.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to get password hash: %v", err,
			)
		}
	case pb.GetPasswordHash() != "":
		passwordHash = []byte(pb.GetPasswordHash())
	}

	var r string

	switch pb.GetRole() {
	case userv2.Role_ROLE_BLUE:
		r = user.Blue
	case userv2.Role_ROLE_RED:
		r = user.Red
	case userv2.Role_ROLE_BLACK:
		r = user.Black
	case userv2.Role_ROLE_UNSPECIFIED:
		r = ""
	}

	return &user.User{
		ID:           id,
		Username:     pb.Username,
		PasswordHash: string(passwordHash),
		TeamID:       teamID,
		Role:         r,
	}, nil
}

func ConvertUserToUserV2Pb(obj *user.User) *userv2.User {
	return &userv2.User{
		Id:           &protov1.UUID{Value: obj.ID.String()},
		Username:     obj.Username,
		TeamId:       &protov1.UUID{Value: obj.TeamID.String()},
		Role:         UserRoleToRoleV2PB(obj.Role),
		PasswordHash: obj.PasswordHash,
	}
}

func UserRoleToRoleV2PB(r string) userv2.Role {
	var rpb userv2.Role
	switch {
	case r == user.Blue:
		rpb = userv2.Role_ROLE_BLUE
	case r == user.Red:
		rpb = userv2.Role_ROLE_RED
	case r == user.Black:
		rpb = userv2.Role_ROLE_BLACK
	}
	return rpb
}
