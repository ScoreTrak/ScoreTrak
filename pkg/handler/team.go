package handler

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamservice"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes/wrappers"
	utilv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	teamv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type TeamController struct {
	svc teamservice.Serv
	teamv1.UnimplementedTeamServiceServer
}

func (p TeamController) GetByID(ctx context.Context, request *teamv1.GetByIDRequest) (*teamv1.GetByIDResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	tm, err := p.svc.GetByID(ctx, uid)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &teamv1.GetByIDResponse{Team: ConvertTeamToTeamPb(tm)}, nil
}

func (p TeamController) GetAll(ctx context.Context, _ *teamv1.GetAllRequest) (*teamv1.GetAllResponse, error) {
	tms, err := p.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	tmspb := make([]*teamv1.Team, 0, len(tms))
	for i := range tms {
		tmspb = append(tmspb, ConvertTeamToTeamPb(tms[i]))
	}
	return &teamv1.GetAllResponse{Teams: tmspb}, nil
}

func (p TeamController) Delete(ctx context.Context, request *teamv1.DeleteRequest) (*teamv1.DeleteResponse, error) {
	uid, err := extractUUID(request)
	if err != nil {
		return nil, err
	}
	err = p.svc.Delete(ctx, uid)
	if err != nil {
		return nil, deleteErrorParser(err)
	}
	return &teamv1.DeleteResponse{}, nil
}

func (p TeamController) Store(ctx context.Context, request *teamv1.StoreRequest) (*teamv1.StoreResponse, error) {
	tmspb := request.GetTeams()
	tms := make([]*team.Team, 0, len(tmspb))
	for i := range tmspb {
		tm, err := ConvertTeamPBtoTeam(false, tmspb[i])
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
	return &teamv1.StoreResponse{Ids: ids}, nil
}

func (p TeamController) Update(ctx context.Context, request *teamv1.UpdateRequest) (*teamv1.UpdateResponse, error) {
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
	return &teamv1.UpdateResponse{}, nil
}

func NewTeamController(svc teamservice.Serv) *TeamController {
	return &TeamController{svc: svc}
}

func ConvertTeamPBtoTeam(requireID bool, pb *teamv1.Team) (*team.Team, error) {
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

func ConvertTeamToTeamPb(obj *team.Team) *teamv1.Team {
	var idx uint64
	if obj.Index != nil {
		idx = *obj.Index
	}
	return &teamv1.Team{
		Id:    &utilv1.UUID{Value: obj.ID.String()},
		Name:  obj.Name,
		Hide:  &wrappers.BoolValue{Value: *obj.Hide},
		Hosts: nil,
		Users: nil,
		Index: &wrappers.UInt64Value{Value: idx},
		Pause: &wrappers.BoolValue{Value: *obj.Pause},
	}
}
