package handler

import (
	"context"
	roundv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/round/v2"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoundV2Controller struct {
	svc roundservice.Serv
	roundv2.UnimplementedRoundServiceServer
}

func (r RoundV2Controller) GetLastNonElapsingRound(ctx context.Context, _ *roundv2.RoundServiceGetLastNonElapsingRoundRequest) (*roundv2.RoundServiceGetLastNonElapsingRoundResponse, error) {
	rnd, err := r.svc.GetLastNonElapsingRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundv2.RoundServiceGetLastNonElapsingRoundResponse{Round: ConvertRoundToRoundV2Pb(rnd)}, nil
}

func (r RoundV2Controller) GetAll(ctx context.Context, _ *roundv2.RoundServiceGetAllRequest) (*roundv2.RoundServiceGetAllResponse, error) {
	rnds, err := r.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	rndspb := make([]*roundv2.Round, 0, len(rnds))
	for i := range rnds {
		rndspb = append(rndspb, ConvertRoundToRoundV2Pb(rnds[i]))
	}
	return &roundv2.RoundServiceGetAllResponse{Rounds: rndspb}, nil
}

func (r RoundV2Controller) GetByID(ctx context.Context, request *roundv2.RoundServiceGetByIDRequest) (*roundv2.RoundServiceGetByIDResponse, error) {
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
	return &roundv2.RoundServiceGetByIDResponse{Round: ConvertRoundToRoundV2Pb(rnd)}, nil
}

func (r RoundV2Controller) GetLastRound(ctx context.Context, _ *roundv2.RoundServiceGetLastRoundRequest) (*roundv2.RoundServiceGetLastRoundResponse, error) {
	rnd, err := r.svc.GetLastRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundv2.RoundServiceGetLastRoundResponse{Round: ConvertRoundToRoundV2Pb(rnd)}, nil
}

func NewRoundV2Controller(svc roundservice.Serv) *RoundV2Controller {
	return &RoundV2Controller{svc: svc}
}

func ConvertRoundV2PBtoRound(requireID bool, pb *roundv2.Round) (*round.Round, error) {
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

func ConvertRoundToRoundV2Pb(obj *round.Round) *roundv2.Round {
	tss := timestamppb.New(obj.Start)
	var tsf *timestamp.Timestamp
	if obj.Finish != nil {
		tsf = timestamppb.New(*obj.Finish)
	}
	return &roundv2.Round{
		Id:     obj.ID,
		Start:  tss,
		Note:   obj.Note,
		Err:    obj.Err,
		Finish: tsf,
		Checks: nil,
	}
}
