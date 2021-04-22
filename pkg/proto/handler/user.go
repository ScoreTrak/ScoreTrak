package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/user_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userpb"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	svc user_service.Serv
	userpb.UnimplementedUserServiceServer
	policyClient *policy_client.Client
}

func (p UserController) GetByUsername(ctx context.Context, request *userpb.GetByUsernameRequest) (*userpb.GetByUsernameResponse, error) {
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
	return &userpb.GetByUsernameResponse{User: ConvertUserToUserPb(tm)}, nil
}

func (p UserController) GetByID(ctx context.Context, request *userpb.GetByIDRequest) (*userpb.GetByIDResponse, error) {
	id := request.GetId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}
	tm, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &userpb.GetByIDResponse{User: ConvertUserToUserPb(tm)}, nil
}

func (p UserController) GetAll(ctx context.Context, _ *userpb.GetAllRequest) (*userpb.GetAllResponse, error) {
	tms, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var tmspb []*userpb.User
	for i := range tms {
		tmspb = append(tmspb, ConvertUserToUserPb(tms[i]))
	}
	return &userpb.GetAllResponse{Users: tmspb}, nil
}

func (p UserController) Delete(ctx context.Context, request *userpb.DeleteRequest) (*userpb.DeleteResponse, error) {
	id := request.GetId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			idNotSpecified,
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			unableToParseID+": %v", err,
		)
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &userpb.DeleteResponse{}, nil
}

func (p UserController) Store(ctx context.Context, request *userpb.StoreRequest) (*userpb.StoreResponse, error) {
	usrspb := request.GetUsers()
	var usrs []*user.User
	for i := range usrspb {
		usr, err := ConvertUserPBtoUser(false, usrspb[i])
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
	err := p.svc.Store(ctx, usrs)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	var ids []*utilpb.UUID
	for i := range usrs {
		ids = append(ids, &utilpb.UUID{Value: usrs[i].ID.String()})
	}
	return &userpb.StoreResponse{Ids: ids}, nil
}

func (p UserController) Update(ctx context.Context, request *userpb.UpdateRequest) (*userpb.UpdateResponse, error) {
	usrpb := request.GetUser()
	usr, err := ConvertUserPBtoUser(true, usrpb)
	if err != nil {
		return nil, err
	}

	claim := extractUserClaim(ctx)

	if claim.Role != user.Black {
		if claim.Id != usr.TeamID.String() {
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

		request.User = &userpb.User{Username: request.GetUser().Username, Password: request.User.GetPassword()}
	}

	err = p.svc.Update(ctx, usr)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &userpb.UpdateResponse{}, nil
}

func NewUserController(svc user_service.Serv, policyClient *policy_client.Client) *UserController {
	return &UserController{svc: svc, policyClient: policyClient}
}

func ConvertUserPBtoUser(requireID bool, pb *userpb.User) (*user.User, error) {
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

	if pb.GetPassword() != "" && pb.GetPasswordHash() != "" {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"You should provide either password or hash, but not both",
		)
	} else if pb.GetPassword() != "" {
		passwordHash, err = bcrypt.GenerateFromPassword([]byte(pb.GetPassword()), bcrypt.DefaultCost)
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to get password hash: %v", err,
			)
		}
	} else if pb.GetPasswordHash() != "" {
		passwordHash = []byte(pb.GetPasswordHash())
	}

	var r string

	if pb.GetRole() == userpb.Role_Blue {
		r = user.Blue
	} else if pb.GetRole() == userpb.Role_Red {
		r = user.Red
	} else if pb.GetRole() == userpb.Role_Black {
		r = user.Black
	}

	return &user.User{
		ID:           id,
		Username:     pb.Username,
		PasswordHash: string(passwordHash),
		TeamID:       teamID,
		Role:         r,
	}, nil
}

func ConvertUserToUserPb(obj *user.User) *userpb.User {
	return &userpb.User{
		Id:           &utilpb.UUID{Value: obj.ID.String()},
		Username:     obj.Username,
		TeamId:       &utilpb.UUID{Value: obj.TeamID.String()},
		Role:         UserRoleToRolePB(obj.Role),
		PasswordHash: obj.PasswordHash,
	}
}

func UserRoleToRolePB(r string) userpb.Role {
	var rpb userpb.Role
	if r == user.Blue {
		rpb = userpb.Role_Blue
	} else if r == user.Red {
		rpb = userpb.Role_Red
	} else if r == user.Black {
		rpb = userpb.Role_Black
	}
	return rpb
}
