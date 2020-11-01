package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/service"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportController struct {
	svc service.Serv
}

func (p ReportController) Get(ctx context.Context, request *reportpb.GetRequest) (*reportpb.GetResponse, error) {
	tms, err := p.svc.Get(ctx)
	if err != nil {
		return nil, getErrorParser(err)
	}

	uat, err := ptypes.TimestampProto(tms.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable convert time.date to timestamp(Ideally, this should not happen, perhaps this is a bug): %v", err),
		)
	}

	return &reportpb.GetResponse{Report: &reportpb.Report{
		Cache:     tms.Cache,
		UpdatedAt: uat,
	}}, nil
}

func NewReportController(svc service.Serv) *ReportController {
	return &ReportController{svc}
}
