package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/service"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RoundController struct {
	svc service.Serv
}

func (r RoundController) GetLastNonElapsingRound(ctx context.Context, request *roundpb.GetLastNonElapsingRoundRequest) (*roundpb.GetLastNonElapsingRoundResponse, error) {
	rnd, err := r.svc.GetLastNonElapsingRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	rndpb, err := ConvertRoundToRoundPb(rnd)
	if err != nil {
		return nil, err
	}
	return &roundpb.GetLastNonElapsingRoundResponse{Round: rndpb}, nil
}

func (r RoundController) GetAll(ctx context.Context, request *roundpb.GetAllRequest) (*roundpb.GetAllResponse, error) {
	rnds, err := r.svc.GetAll(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	var rndspb []*roundpb.Round
	for i := range rnds {
		rndpb, err := ConvertRoundToRoundPb(rnds[i])
		if err != nil {
			return nil, err
		}
		rndspb = append(rndspb, rndpb)
	}
	return &roundpb.GetAllResponse{Rounds: rndspb}, nil
}

func (r RoundController) GetByID(ctx context.Context, request *roundpb.GetByIDRequest) (*roundpb.GetByIDResponse, error) {
	roundID := request.GetId()
	if roundID == 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"Round ID was not specified",
		)
	}

	rnd, err := r.svc.GetByID(ctx, roundID)
	if err != nil {
		return nil, getErrorParser(err)
	}
	rndpb, err := ConvertRoundToRoundPb(rnd)
	if err != nil {
		return nil, err
	}
	return &roundpb.GetByIDResponse{Round: rndpb}, nil
}

func (r RoundController) GetLastRound(ctx context.Context, request *roundpb.GetLastRoundRequest) (*roundpb.GetLastRoundResponse, error) {
	rnd, err := r.svc.GetLastRound(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}
	rndpb, err := ConvertRoundToRoundPb(rnd)
	if err != nil {
		return nil, err
	}
	return &roundpb.GetLastRoundResponse{Round: rndpb}, nil
}

func NewRoundController(svc service.Serv) *RoundController {
	return &RoundController{svc}
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

func ConvertRoundToRoundPb(obj *round.Round) (*roundpb.Round, error) {
	tss, err := ptypes.TimestampProto(obj.Start)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable convert time.date to timestamp(Ideally this shouldn't be happening, and this is most likely a bug): %v", err),
		)
	}

	var tsf *timestamp.Timestamp
	if obj.Finish != nil {
		tsf, err = ptypes.TimestampProto(*obj.Finish)
		if err != nil {
			return nil, status.Errorf(
				codes.Internal,
				fmt.Sprintf("Unable convert time.date to timestamp(Ideally this shouldn't be happening, and this is most likely a bug): %v", err),
			)
		}
	}

	return &roundpb.Round{
		Id:     obj.ID,
		Start:  tss,
		Note:   obj.Note,
		Err:    obj.Err,
		Finish: tsf,
		Checks: nil,
	}, nil
}
