package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	roundpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/round/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoundController struct {
	svc roundservice.Serv
	roundpb.UnimplementedRoundServiceServer
}

func (r RoundController) GetLastNonElapsingRound(ctx context.Context, _ *roundpb.GetLastNonElapsingRoundRequest) (*roundpb.GetLastNonElapsingRoundResponse, error) {
	rnd, err := r.svc.GetLastNonElapsingRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundpb.GetLastNonElapsingRoundResponse{Round: ConvertRoundToRoundPb(rnd)}, nil
}

func (r RoundController) GetAll(ctx context.Context, _ *roundpb.GetAllRequest) (*roundpb.GetAllResponse, error) {
	rnds, err := r.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	rndspb := make([]*roundpb.Round, 0, len(rnds))
	for i := range rnds {
		rndspb = append(rndspb, ConvertRoundToRoundPb(rnds[i]))
	}
	return &roundpb.GetAllResponse{Rounds: rndspb}, nil
}

func (r RoundController) GetByID(ctx context.Context, request *roundpb.GetByIDRequest) (*roundpb.GetByIDResponse, error) {
	roundID := request.GetId()
	if roundID == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round"+idNotSpecified,
		)
	}

	rnd, err := r.svc.GetByID(ctx, roundID)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundpb.GetByIDResponse{Round: ConvertRoundToRoundPb(rnd)}, nil
}

func (r RoundController) GetLastRound(ctx context.Context, _ *roundpb.GetLastRoundRequest) (*roundpb.GetLastRoundResponse, error) {
	rnd, err := r.svc.GetLastRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundpb.GetLastRoundResponse{Round: ConvertRoundToRoundPb(rnd)}, nil
}

func NewRoundController(svc roundservice.Serv) *RoundController {
	return &RoundController{svc: svc}
}

func ConvertRoundPBtoRound(requireID bool, pb *roundpb.Round) (*round.Round, error) {
	if pb.Id == 0 && requireID {
		return nil, status.Error(
			codes.InvalidArgument,
			"ID is required",
		)
	}
	tfs := pb.Finish.AsTime()
	return &round.Round{
		ID:     pb.Id,
		Start:  pb.Start.AsTime(),
		Note:   pb.Note,
		Err:    pb.Err,
		Finish: &tfs,
		Checks: nil,
	}, nil
}

func ConvertRoundToRoundPb(obj *round.Round) *roundpb.Round {
	tss := timestamppb.New(obj.Start)
	var tsf *timestamp.Timestamp
	if obj.Finish != nil {
		tsf = timestamppb.New(*obj.Finish)
	}
	return &roundpb.Round{
		Id:     obj.ID,
		Start:  tss,
		Note:   obj.Note,
		Err:    obj.Err,
		Finish: tsf,
		Checks: nil,
	}
}
