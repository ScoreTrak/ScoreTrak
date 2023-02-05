package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/policy/v2/policyv2grpc"
	policyv2 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/policy/v2"
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PolicyV2Controller struct {
	svc          policyservice.Serv
	policyClient *policyclient.Client
	policyv2grpc.UnimplementedPolicyServiceServer
}

func (p PolicyV2Controller) Get(_ *policyv2.PolicyServiceGetRequest, server policyv2grpc.PolicyService_GetServer) error {
	rol := user.Anonymous
	claims := extractUserClaim(server.Context())
	if claims != nil {
		rol = claims.Role
	}

	err := server.Send(&policyv2.PolicyServiceGetResponse{
		Policy: ConvertPolicyToPolicyV2PB(p.policyClient.GetPolicy()),
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
			err := server.Send(&policyv2.PolicyServiceGetResponse{
				Policy: ConvertPolicyToPolicyV2PB(p.policyClient.GetPolicy()),
			})
			if err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

func (p PolicyV2Controller) GetUnary(ctx context.Context, request *policyv2.PolicyServiceGetUnaryRequest) (*policyv2.PolicyServiceGetUnaryResponse, error) {
	pol, err := p.svc.Get(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &policyv2.PolicyServiceGetUnaryResponse{Policy: ConvertPolicyToPolicyV2PB(pol)}, nil
}

func (p PolicyV2Controller) Update(ctx context.Context, request *policyv2.PolicyServiceUpdateRequest) (*policyv2.PolicyServiceUpdateResponse, error) {
	polpb := request.GetPolicy()
	err := p.svc.Update(ctx, ConvertPolicyV2PBToPolicy(polpb))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	p.policyClient.Notify()
	return &policyv2.PolicyServiceUpdateResponse{}, nil
}

func NewPolicyV2Controller(svc policyservice.Serv, client *policyclient.Client) *PolicyV2Controller {
	return &PolicyV2Controller{
		svc:          svc,
		policyClient: client,
	}
}

func ConvertPolicyV2PBToPolicy(pb *policyv2.Policy) *policy.Policy {
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

func ConvertPolicyToPolicyV2PB(obj *policy.Policy) *policyv2.Policy {
	return &policyv2.Policy{
		AllowUnauthenticatedUsers:                 &wrappers.BoolValue{Value: *obj.AllowUnauthenticatedUsers},
		AllowChangingUsernamesAndPasswords:        &wrappers.BoolValue{Value: *obj.AllowChangingUsernamesAndPasswords},
		ShowPoints:                                &wrappers.BoolValue{Value: *obj.ShowPoints},
		ShowAddresses:                             &wrappers.BoolValue{Value: *obj.ShowAddresses},
		AllowRedTeamLaunchingServiceTestsManually: &wrappers.BoolValue{Value: *obj.AllowRedTeamLaunchingServiceTestsManually},
	}
}
