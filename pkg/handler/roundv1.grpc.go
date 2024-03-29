package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/round/v1/roundv1grpc"
	roundv1 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/round/v1"
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoundV1Controller struct {
	svc roundservice.Serv
	roundv1grpc.UnimplementedRoundServiceServer
}

func (r RoundV1Controller) GetLastNonElapsingRound(ctx context.Context, _ *roundv1.GetLastNonElapsingRoundRequest) (*roundv1.GetLastNonElapsingRoundResponse, error) {
	rnd, err := r.svc.GetLastNonElapsingRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundv1.GetLastNonElapsingRoundResponse{Round: ConvertRoundToRoundV1Pb(rnd)}, nil
}

func (r RoundV1Controller) GetAll(ctx context.Context, _ *roundv1.GetAllRequest) (*roundv1.GetAllResponse, error) {
	rnds, err := r.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	rndspb := make([]*roundv1.Round, 0, len(rnds))
	for i := range rnds {
		rndspb = append(rndspb, ConvertRoundToRoundV1Pb(rnds[i]))
	}
	return &roundv1.GetAllResponse{Rounds: rndspb}, nil
}

func (r RoundV1Controller) GetByID(ctx context.Context, request *roundv1.GetByIDRequest) (*roundv1.GetByIDResponse, error) {
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
	return &roundv1.GetByIDResponse{Round: ConvertRoundToRoundV1Pb(rnd)}, nil
}

func (r RoundV1Controller) GetLastRound(ctx context.Context, _ *roundv1.GetLastRoundRequest) (*roundv1.GetLastRoundResponse, error) {
	rnd, err := r.svc.GetLastRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	return &roundv1.GetLastRoundResponse{Round: ConvertRoundToRoundV1Pb(rnd)}, nil
}

func NewRoundV1Controller(svc roundservice.Serv) *RoundV1Controller {
	return &RoundV1Controller{svc: svc}
}

func ConvertRoundV1PBtoRound(requireID bool, pb *roundv1.Round) (*round.Round, error) {
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

func ConvertRoundToRoundV1Pb(obj *round.Round) *roundv1.Round {
	tss := timestamppb.New(obj.Start)
	var tsf *timestamp.Timestamp
	if obj.Finish != nil {
		tsf = timestamppb.New(*obj.Finish)
	}
	return &roundv1.Round{
		Id:     obj.ID,
		Start:  tss,
		Note:   obj.Note,
		Err:    obj.Err,
		Finish: tsf,
		Checks: nil,
	}
}
