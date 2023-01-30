package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	utilv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	teamv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeamV2Controller struct {
	svc teamservice.Serv
	teamv2.UnimplementedTeamServiceServer
}

func (p TeamV2Controller) GetByID(ctx context.Context, request *teamv2.TeamServiceGetByIDRequest) (*teamv2.TeamServiceGetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	tm, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &teamv2.TeamServiceGetByIDResponse{Team: ConvertTeamToTeamV2Pb(tm)}, nil
}

func (p TeamV2Controller) GetAll(ctx context.Context, _ *teamv2.TeamServiceGetAllRequest) (*teamv2.TeamServiceGetAllResponse, error) {
	tms, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	tmspb := make([]*teamv2.Team, 0, len(tms))
	for i := range tms {
		tmspb = append(tmspb, ConvertTeamToTeamV2Pb(tms[i]))
	}
	return &teamv2.TeamServiceGetAllResponse{Teams: tmspb}, nil
}

func (p TeamV2Controller) Delete(ctx context.Context, request *teamv2.TeamServiceDeleteRequest) (*teamv2.TeamServiceDeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &teamv2.TeamServiceDeleteResponse{}, nil
}

func (p TeamV2Controller) Store(ctx context.Context, request *teamv2.TeamServiceStoreRequest) (*teamv2.TeamServiceStoreResponse, error) {
	tmspb := request.GetTeams()
	tms := make([]*team.Team, 0, len(tmspb))
	for i := range tmspb {
		tm, err := ConvertTeamV2PBtoTeam(false, tmspb[i])
		if err != nil {
			return nil, err
		}
		tms = append(tms, tm)
	}
	if err := p.svc.Store(ctx, tms); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	ids := make([]*utilv1.UUID, 0, len(tms))
	for i := range tms {
		ids = append(ids, &utilv1.UUID{Value: tms[i].ID.String()})
	}
	return &teamv2.TeamServiceStoreResponse{Ids: ids}, nil
}

func (p TeamV2Controller) Update(ctx context.Context, request *teamv2.TeamServiceUpdateRequest) (*teamv2.TeamServiceUpdateResponse, error) {
	tmspb := request.GetTeam()
	tm, err := ConvertTeamV2PBtoTeam(true, tmspb)
	if err != nil {
		return nil, err
	}
	err = p.svc.Update(ctx, tm)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	return &teamv2.TeamServiceUpdateResponse{}, nil
}

func NewTeamV2Controller(svc teamservice.Serv) *TeamV2Controller {
	return &TeamV2Controller{svc: svc}
}

func ConvertTeamV2PBtoTeam(requireID bool, pb *teamv2.Team) (*team.Team, error) {
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
	var pause *bool
	if pb.GetPause() != nil {
		pause = &pb.GetPause().Value
	}

	var hide *bool
	if pb.GetHide() != nil {
		hide = &pb.GetHide().Value
	}

	var index *uint64
	if pb.GetIndex() != nil {
		index = &pb.GetIndex().Value
	}

	return &team.Team{
		ID:    id,
		Name:  pb.GetName(),
		Pause: pause,
		Users: nil,
		Hosts: nil,
		Index: index,
		Hide:  hide,
	}, nil
}

func ConvertTeamToTeamV2Pb(obj *team.Team) *teamv2.Team {
	var idx uint64
	if obj.Index != nil {
		idx = *obj.Index
	}
	return &teamv2.Team{
		Id:    &utilv1.UUID{Value: obj.ID.String()},
		Name:  obj.Name,
		Hide:  &wrappers.BoolValue{Value: *obj.Hide},
		Hosts: nil,
		Users: nil,
		Index: &wrappers.UInt64Value{Value: idx},
		Pause: &wrappers.BoolValue{Value: *obj.Pause},
	}
}
