package handler

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyService"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	policypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/policy/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PolicyController struct {
	svc          policyService.Serv
	policyClient *policy_client.Client
	policypb.UnimplementedPolicyServiceServer
}

func (p PolicyController) Get(request *policypb.GetRequest, server policypb.PolicyService_GetServer) error {
	rol := user.Anonymous
	claims := extractUserClaim(server.Context())
	if claims != nil {
		rol = claims.Role
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
			if !p.policyClient.GetAllowUnauthenticatedUsers() && rol == user.Anonymous {
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

func NewPolicyController(svc policyService.Serv, client *policy_client.Client) *PolicyController {
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

	var artlstm *bool
	if pb.GetAllowRedTeamLaunchingServiceTestsManually() != nil {
		artlstm = &pb.GetAllowRedTeamLaunchingServiceTestsManually().Value
	}

	return &policy.Policy{
		AllowUnauthenticatedUsers:                 auu,
		AllowChangingUsernamesAndPasswords:        acup,
		AllowRedTeamLaunchingServiceTestsManually: artlstm,
		ShowPoints:    sp,
		ShowAddresses: sa,
	}
}

func ConvertPolicyToPolicyPB(obj *policy.Policy) *policypb.Policy {
	return &policypb.Policy{
		AllowUnauthenticatedUsers:                 &wrappers.BoolValue{Value: *obj.AllowUnauthenticatedUsers},
		AllowChangingUsernamesAndPasswords:        &wrappers.BoolValue{Value: *obj.AllowChangingUsernamesAndPasswords},
		ShowPoints:                                &wrappers.BoolValue{Value: *obj.ShowPoints},
		ShowAddresses:                             &wrappers.BoolValue{Value: *obj.ShowAddresses},
		AllowRedTeamLaunchingServiceTestsManually: &wrappers.BoolValue{Value: *obj.AllowRedTeamLaunchingServiceTestsManually},
	}
}
