package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/policy/v1/policyv1grpc"
	policyv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/policy/v1"
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

type PolicyV1Controller struct {
	svc          policyservice.Serv
	policyClient *policyclient.Client
	policyv1grpc.UnimplementedPolicyServiceServer
}

func (p PolicyV1Controller) Get(_ *policyv1.GetRequest, server policyv1grpc.PolicyService_GetServer) error {
	rol := user.Anonymous
	claims := extractUserClaim(server.Context())
	if claims != nil {
		rol = claims.Role
	}

	err := server.Send(&policyv1.GetResponse{
		Policy: ConvertPolicyToPolicyV1PB(p.policyClient.GetPolicy()),
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
			err := server.Send(&policyv1.GetResponse{
				Policy: ConvertPolicyToPolicyV1PB(p.policyClient.GetPolicy()),
			})
			if err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

func (p PolicyV1Controller) GetUnary(ctx context.Context, request *policyv1.GetUnaryRequest) (*policyv1.GetUnaryResponse, error) {
	pol, err := p.svc.Get(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &policyv1.GetUnaryResponse{Policy: ConvertPolicyToPolicyV1PB(pol)}, nil
}

func (p PolicyV1Controller) Update(ctx context.Context, request *policyv1.UpdateRequest) (*policyv1.UpdateResponse, error) {
	polpb := request.GetPolicy()
	err := p.svc.Update(ctx, ConvertPolicyV1PBToPolicy(polpb))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	p.policyClient.Notify()
	return &policyv1.UpdateResponse{}, nil
}

func NewPolicyV1Controller(svc policyservice.Serv, client *policyclient.Client) *PolicyV1Controller {
	return &PolicyV1Controller{
		svc:          svc,
		policyClient: client,
	}
}

func ConvertPolicyV1PBToPolicy(pb *policyv1.Policy) *policy.Policy {
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

func ConvertPolicyToPolicyV1PB(obj *policy.Policy) *policyv1.Policy {
	return &policyv1.Policy{
		AllowUnauthenticatedUsers:                 &wrappers.BoolValue{Value: *obj.AllowUnauthenticatedUsers},
		AllowChangingUsernamesAndPasswords:        &wrappers.BoolValue{Value: *obj.AllowChangingUsernamesAndPasswords},
		ShowPoints:                                &wrappers.BoolValue{Value: *obj.ShowPoints},
		ShowAddresses:                             &wrappers.BoolValue{Value: *obj.ShowAddresses},
		AllowRedTeamLaunchingServiceTestsManually: &wrappers.BoolValue{Value: *obj.AllowRedTeamLaunchingServiceTestsManually},
	}
}
