package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/team_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teampb"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeamController struct {
	svc team_service.Serv
}

func (p TeamController) GetByID(ctx context.Context, request *teampb.GetByIDRequest) (*teampb.GetByIDResponse, error) {
	id := request.GetId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	tm, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &teampb.GetByIDResponse{Team: ConvertTeamToTeamPb(tm)}, nil
}

func (p TeamController) GetAll(ctx context.Context, request *teampb.GetAllRequest) (*teampb.GetAllResponse, error) {
	tms, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var tmspb []*teampb.Team
	for i := range tms {
		tmspb = append(tmspb, ConvertTeamToTeamPb(tms[i]))
	}
	return &teampb.GetAllResponse{Teams: tmspb}, nil
}

func (p TeamController) Delete(ctx context.Context, request *teampb.DeleteRequest) (*teampb.DeleteResponse, error) {
	id := request.GetId()
	if id == nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	uid, err := uuid.FromString(id.GetValue())
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Unable to parse ID: %v", err,
		)
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &teampb.DeleteResponse{}, nil
}

func (p TeamController) Store(ctx context.Context, request *teampb.StoreRequest) (*teampb.StoreResponse, error) {
	tmspb := request.GetTeams()
	var tms []*team.Team
	for i := range tmspb {
		tm, err := ConvertTeamPBtoTeam(false, tmspb[i])
		if err != nil {
			return nil, err
		}
		tms = append(tms, tm)
	}
	err := p.svc.Store(ctx, tms)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unknown internal error: %v", err),
		)
	}
	var ids []*utilpb.UUID
	for i := range tms {
		ids = append(ids, &utilpb.UUID{Value: tms[i].ID.String()})
	}
	return &teampb.StoreResponse{Ids: ids}, nil
}

func (p TeamController) Update(ctx context.Context, request *teampb.UpdateRequest) (*teampb.UpdateResponse, error) {
	tmspb := request.GetTeam()
	tm, err := ConvertTeamPBtoTeam(true, tmspb)
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
	return &teampb.UpdateResponse{}, nil
}

func NewTeamController(svc team_service.Serv) *TeamController {
	return &TeamController{svc}
}

func ConvertTeamPBtoTeam(requireID bool, pb *teampb.Team) (*team.Team, error) {
	var id uuid.UUID
	var err error
	if pb.GetId() != nil {
		id, err = uuid.FromString(pb.GetId().GetValue())
		if err != nil {
			return nil, status.Errorf(
				codes.InvalidArgument,
				"Unable to parse ID: %v", err,
			)
		}
	} else if requireID {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"ID was not specified",
		)
	}
	var enabled *bool
	if pb.GetEnabled() != nil {
		enabled = &pb.GetEnabled().Value
	}

	var hidden *bool
	if pb.GetHidden() != nil {
		hidden = &pb.GetHidden().Value
	}

	var index *uint64
	if pb.GetIndex() != nil {
		index = &pb.GetIndex().Value
	}

	return &team.Team{
		ID:      id,
		Name:    pb.GetName(),
		Enabled: enabled,
		Users:   nil,
		Hosts:   nil,
		Index:   index,
		Hidden:  hidden,
	}, nil
}

func ConvertTeamToTeamPb(obj *team.Team) *teampb.Team {
	var idx uint64
	if obj.Index != nil {
		idx = *obj.Index
	}
	return &teampb.Team{
		Id:      &utilpb.UUID{Value: obj.ID.String()},
		Name:    obj.Name,
		Enabled: &wrappers.BoolValue{Value: *obj.Enabled},
		Hosts:   nil,
		Users:   nil,
		Index:   &wrappers.UInt64Value{Value: idx},
		Hidden:  &wrappers.BoolValue{Value: *obj.Hidden},
	}
}
