package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/user/v1/userv1grpc"
	protov1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/proto/v1"
	userv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/user/v1"
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserV1Controller struct {
	svc userservice.Serv
	userv1grpc.UnimplementedUserServiceServer
	policyClient *policyclient.Client
}

func (p UserV1Controller) GetByUsername(ctx context.Context, request *userv1.GetByUsernameRequest) (*userv1.GetByUsernameResponse, error) {
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
	return &userv1.GetByUsernameResponse{User: ConvertUserToUserV1Pb(tm)}, nil
}

func (p UserV1Controller) GetByID(ctx context.Context, request *userv1.GetByIDRequest) (*userv1.GetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	tm, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &userv1.GetByIDResponse{User: ConvertUserToUserV1Pb(tm)}, nil
}

func (p UserV1Controller) GetAll(ctx context.Context, _ *userv1.GetAllRequest) (*userv1.GetAllResponse, error) {
	tms, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	tmspb := make([]*userv1.User, 0, len(tms))
	for i := range tms {
		tmspb = append(tmspb, ConvertUserToUserV1Pb(tms[i]))
	}
	return &userv1.GetAllResponse{Users: tmspb}, nil
}

func (p UserV1Controller) Delete(ctx context.Context, request *userv1.DeleteRequest) (*userv1.DeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &userv1.DeleteResponse{}, nil
}

func (p UserV1Controller) Store(ctx context.Context, request *userv1.StoreRequest) (*userv1.StoreResponse, error) {
	usrspb := request.GetUsers()
	usrs := make([]*user.User, 0, len(usrspb))
	for i := range usrspb {
		usr, err := ConvertUserV1PBtoUser(false, usrspb[i])
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
	return &userv1.StoreResponse{Ids: ids}, nil
}

func (p UserV1Controller) Update(ctx context.Context, request *userv1.UpdateRequest) (*userv1.UpdateResponse, error) {
	usrpb := request.GetUser()
	usr, err := ConvertUserV1PBtoUser(true, usrpb)
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

		request.User = &userv1.User{Username: request.GetUser().Username, Password: request.User.GetPassword()}
	}

	err = p.svc.Update(ctx, usr)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &userv1.UpdateResponse{}, nil
}

func NewUserV1Controller(svc userservice.Serv, policyClient *policyclient.Client) *UserV1Controller {
	return &UserV1Controller{svc: svc, policyClient: policyClient}
}

func ConvertUserV1PBtoUser(requireID bool, pb *userv1.User) (*user.User, error) {
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
	case userv1.Role_ROLE_BLUE:
		r = user.Blue
	case userv1.Role_ROLE_RED:
		r = user.Red
	case userv1.Role_ROLE_BLACK:
		r = user.Black
	case userv1.Role_ROLE_UNSPECIFIED:
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

func ConvertUserToUserV1Pb(obj *user.User) *userv1.User {
	return &userv1.User{
		Id:           &protov1.UUID{Value: obj.ID.String()},
		Username:     obj.Username,
		TeamId:       &protov1.UUID{Value: obj.TeamID.String()},
		Role:         UserRoleToRoleV1PB(obj.Role),
		PasswordHash: obj.PasswordHash,
	}
}

func UserRoleToRoleV1PB(r string) userv1.Role {
	var rpb userv1.Role
	switch {
	case r == user.Blue:
		rpb = userv1.Role_ROLE_BLUE
	case r == user.Red:
		rpb = userv1.Role_ROLE_RED
	case r == user.Black:
		rpb = userv1.Role_ROLE_BLACK
	}
	return rpb
}
