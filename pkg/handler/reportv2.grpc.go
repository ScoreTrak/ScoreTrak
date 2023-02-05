package handler

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/report/v2/reportv2grpc"
	reportv2 "buf.build/gen/go/scoretrak/scoretrakapis/protocolbuffers/go/scoretrak/report/v2"
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportV2Controller struct {
	svc          reportservice.Serv
	reportClient *reportclient.Client
	policyClient *policyclient.Client
	reportv2grpc.UnimplementedReportServiceServer
}

func (r *ReportV2Controller) filterReport(rol string, tID uuid.UUID, lr *report.Report) (*reportv2.Report, error) {
	simpleReport := &report.SimpleReport{}
	err := json.Unmarshal([]byte(lr.Cache), simpleReport)
	if err != nil {
		return nil, err
	}
	removeDisabledAndHidden(simpleReport)
	calculateTotalPoints(simpleReport)
	if rol != user.Black {
		filterBlueTeams(simpleReport, tID, r.policyClient.GetPolicy())
	}
	ret, err := json.Marshal(simpleReport)
	if err != nil {
		return nil, err
	}
	return &reportv2.Report{
		Cache:     string(ret),
		UpdatedAt: timestamppb.New(lr.UpdatedAt),
	}, nil
}

func (r *ReportV2Controller) GetUnary(ctx context.Context, _ *reportv2.ReportServiceGetUnaryRequest) (*reportv2.ReportServiceGetUnaryResponse, error) {
	rol := user.Anonymous
	tID := uuid.UUID{}
	lr, err := r.svc.Get(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			fmt.Sprintf("Unable to retrieve report: %v", err))
	}

	claims := extractUserClaim(ctx)
	if claims != nil {
		rol = claims.Role
		tID = uuid.FromStringOrNil(claims.TeamID)
	}

	if rol == user.Anonymous && !r.policyClient.GetAllowUnauthenticatedUsers() {
		return nil, status.Error(codes.PermissionDenied, "You must login in order to access this resource")
	}

	frep, err := r.filterReport(rol, tID, lr)
	if err != nil {
		return nil, status.Errorf(codes.Internal,
			fmt.Sprintf("Unable to filter report: %v", err))
	}

	return &reportv2.ReportServiceGetUnaryResponse{Report: frep}, nil
}

func (r *ReportV2Controller) Get(_ *reportv2.ReportServiceGetRequest, server reportv2grpc.ReportService_GetServer) error {
	rol := user.Anonymous
	tID := uuid.UUID{}
	lr, err := r.svc.Get(server.Context())
	if err != nil {
		return status.Errorf(codes.Internal,
			fmt.Sprintf("Unable to retrieve report: %v", err))
	}

	claims := extractUserClaim(server.Context())
	if claims != nil {
		rol = claims.Role
		tID = uuid.FromStringOrNil(claims.TeamID)
	}

	if rol == user.Anonymous && !r.policyClient.GetAllowUnauthenticatedUsers() {
		return status.Error(codes.PermissionDenied, "You must login in order to access this resource")
	}

	frep, err := r.filterReport(rol, tID, lr)
	if err != nil {
		return status.Errorf(codes.Internal,
			fmt.Sprintf("Unable to filter report: %v", err))
	}
	err = server.Send(&reportv2.ReportServiceGetResponse{Report: frep})
	if err != nil {
		return err
	}

	uid, ch := r.reportClient.Subscribe()
	defer r.reportClient.Unsubscribe(uid)

	for {
		select {
		case <-ch:
			if rol == user.Anonymous && !r.policyClient.GetAllowUnauthenticatedUsers() {
				return status.Error(codes.PermissionDenied, "You must login in order to access this resource")
			}
			lr, err := r.svc.Get(server.Context())
			if err != nil {
				return status.Errorf(codes.Internal,
					fmt.Sprintf("Unable to retrieve report: %v", err))
			}
			frep, err = r.filterReport(rol, tID, lr)
			if err != nil {
				return status.Errorf(codes.Internal,
					fmt.Sprintf("Unable to filter report: %v", err))
			}
			err = server.Send(&reportv2.ReportServiceGetResponse{Report: frep})
			if err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

func NewReportV2Controller(svc reportservice.Serv, reportClient *reportclient.Client, client *policyclient.Client) *ReportV2Controller {
	return &ReportV2Controller{
		svc:          svc,
		reportClient: reportClient,
		policyClient: client,
	}
}
