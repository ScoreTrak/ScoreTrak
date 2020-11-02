package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policypb"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/service"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PolicyController struct {
	svc service.Serv
}

func (p PolicyController) Get(ctx context.Context, request *policypb.GetRequest) (*policypb.GetResponse, error) {
	pol, err := p.svc.Get(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &policypb.GetResponse{Policy: ConvertPolicyToPolicyPB(pol)}, nil
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
	return &policypb.UpdateResponse{}, nil
}

func NewPolicyController(svc service.Serv) *PolicyController {
	return &PolicyController{svc}
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
