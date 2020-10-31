package handler

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/logger"
)

type checkController struct {
	log logger.LogInfoFormat
	svc check.Serv
}

func (c checkController) GetAllByRoundID(ctx context.Context, request *checkpb.GetAllByRoundIDRequest) (*checkpb.GetAllByRoundIDResponse, error) {
	panic("implement me")
}

func (c checkController) GetByRoundServiceID(ctx context.Context, request *checkpb.GetByRoundServiceIDRequest) (*checkpb.GetByRoundServiceIDResponse, error) {
	panic("implement me")
}

func (c checkController) GetAllByServiceID(ctx context.Context, request *checkpb.GetAllByServiceIDRequest) (*checkpb.GetAllByServiceIDResponse, error) {
	panic("implement me")
}

func NewCheckController(log logger.LogInfoFormat, svc check.Serv) *checkController {
	return &checkController{log, svc}
}
