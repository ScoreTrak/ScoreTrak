package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policypb"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PolicyController struct {
	svc          policy_service.Serv
	policyClient *policy_client.Client
}

func (p PolicyController) Get(request *policypb.GetRequest, server policypb.PolicyService_GetServer) error {
	rol := role.Anonymous
	if val, ok := server.Context().Value("claims").(*auth.UserClaims); ok && val != nil {
		rol = val.Role
	}

	err := server.Send(&policypb.GetResponse{
		Policy: ConvertPolicyToPolicyPB(p.policyClient.GetPolicy()),
	})
	if err != nil {
		return err
	}
	uuid, ch := p.policyClient.Subscribe()

	defer p.policyClient.Unsubscribe(uuid)
	for {
		select {
		case <-ch:
			if !p.policyClient.GetAllowUnauthenticatedUsers() && rol == role.Anonymous {
				return status.Error(codes.PermissionDenied, "You must login in order to access this resource")
			}
			err := server.Send(&policypb.GetResponse{
				Policy: ConvertPolicyToPolicyPB(p.policyClient.GetPolicy()),
			})
			if err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

func (p PolicyController) Update(ctx context.Context, request *policypb.UpdateRequest) (*policypb.UpdateResponse, error) {
	polpb := request.GetPolicy()
	err := p.svc.Update(ctx, ConvertPolicyPBToPolicy(polpb))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	p.policyClient.Notify()
	return &policypb.UpdateResponse{}, nil
}

func NewPolicyController(svc policy_service.Serv, client *policy_client.Client) *PolicyController {
	return &PolicyController{
		svc:          svc,
		policyClient: client,
	}
}

func ConvertPolicyPBToPolicy(pb *policypb.Policy) *policy.Policy {
	var auu *bool
	if pb.GetAllowUnauthenticatedUsers() != nil {
		auu = &pb.GetAllowUnauthenticatedUsers().Value
	}

	var acup *bool
	if pb.GetAllowChangingUsernamesAndPasswords() != nil {
		acup = &pb.GetAllowChangingUsernamesAndPasswords().Value
	}

	var sp *bool
	if pb.GetShowPoints() != nil {
		sp = &pb.GetShowPoints().Value
	}

	var sa *bool
	if pb.GetShowAddresses() != nil {
		sa = &pb.GetShowAddresses().Value
	}

	return &policy.Policy{
		AllowUnauthenticatedUsers:          auu,
		AllowChangingUsernamesAndPasswords: acup,
		ShowPoints:                         sp,
		ShowAddresses:                      sa,
	}

}

func ConvertPolicyToPolicyPB(obj *policy.Policy) *policypb.Policy {
	return &policypb.Policy{
		AllowUnauthenticatedUsers:          &wrappers.BoolValue{Value: *obj.AllowUnauthenticatedUsers},
		AllowChangingUsernamesAndPasswords: &wrappers.BoolValue{Value: *obj.AllowChangingUsernamesAndPasswords},
		ShowPoints:                         &wrappers.BoolValue{Value: *obj.ShowPoints},
		ShowAddresses:                      &wrappers.BoolValue{Value: *obj.ShowAddresses},
	}
}
