package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	reportpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/report/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_client"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type reportController struct {
	svc          report_service.Serv
	reportClient *report_client.Client
	policyClient *policy_client.Client
	reportpb.UnimplementedReportServiceServer
}

func (r *reportController) filterReport(rol string, tID uuid.UUID, lr *report.Report) (*reportpb.Report, error) {
	simpleReport := &report.SimpleReport{}
	err := json.Unmarshal([]byte(lr.Cache), simpleReport)
	if err != nil {
		return nil, err
	}
	//Filter out Disabled, and Hidden Services
	{
		for t := range simpleReport.Teams {
			if simpleReport.Teams[t].Hide {
				delete(simpleReport.Teams, t)
				continue
			}
			for h := range simpleReport.Teams[t].Hosts {
				if simpleReport.Teams[t].Hosts[h].Hide || (simpleReport.Teams[t].Hosts[h].HostGroup != nil && simpleReport.Teams[t].Hosts[h].HostGroup.Hide) {
					delete(simpleReport.Teams[t].Hosts, h)
					continue
				}
				for s := range simpleReport.Teams[t].Hosts[h].Services {
					if simpleReport.Teams[t].Hosts[h].Services[s].Hide || !simpleReport.Teams[t].Hosts[h].Services[s].SimpleServiceGroup.Enabled {
						delete(simpleReport.Teams[t].Hosts[h].Services, s)
						continue
					}
				}
			}
		}
	}

	//Calculate TotalPoints
	{
		for t := range simpleReport.Teams {
			for h := range simpleReport.Teams[t].Hosts {
				for s := range simpleReport.Teams[t].Hosts[h].Services {
					simpleReport.Teams[t].TotalPoints += simpleReport.Teams[t].Hosts[h].Services[s].Points + simpleReport.Teams[t].Hosts[h].Services[s].PointsBoost
				}
			}
		}
	}

	if rol != user.Black {
		p := r.policyClient.GetPolicy()
		for t := range simpleReport.Teams {
			for h := range simpleReport.Teams[t].Hosts {
				for s := range simpleReport.Teams[t].Hosts[h].Services {
					propFilterHide := map[string]*report.SimpleProperty{}
					for key, val := range simpleReport.Teams[t].Hosts[h].Services[s].Properties {
						if val.Status != "Hide" {
							propFilterHide[key] = val
						}
					}
					simpleReport.Teams[t].Hosts[h].Services[s].Properties = propFilterHide
					if t != tID {
						if simpleReport.Teams[t].Hosts[h].Services[s].Check != nil {
							simpleReport.Teams[t].Hosts[h].Services[s].Check.Err = ""
							simpleReport.Teams[t].Hosts[h].Services[s].Check.Log = ""
						}
						prop := map[string]*report.SimpleProperty{}
						if *p.ShowAddresses {
							if val, ok := simpleReport.Teams[t].Hosts[h].Services[s].Properties["Port"]; ok {
								prop["Port"] = val
							}
						}
						simpleReport.Teams[t].Hosts[h].Services[s].Properties = prop
						if !*p.ShowPoints {
							simpleReport.Teams[t].Hosts[h].Services[s].Points = 0
							simpleReport.Teams[t].Hosts[h].Services[s].PointsBoost = 0
							simpleReport.Teams[t].TotalPoints = 0
						}
					}
				}
				if t != tID {
					if !*p.ShowAddresses {
						simpleReport.Teams[t].Hosts[h].Address = ""
					}
				}
			}
		}
	}
	ret, err := json.Marshal(simpleReport)
	if err != nil {
		return nil, err
	}
	uat, err := ptypes.TimestampProto(lr.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable convert time.date to timestamp(Ideally, this should not happen, perhaps this is a bug): %v", err),
		)
	}
	return &reportpb.Report{
		Cache:     string(ret),
		UpdatedAt: uat,
	}, nil
}

func (r *reportController) Get(request *reportpb.GetRequest, server reportpb.ReportService_GetServer) error {
	rol := user.Anonymous
	tID := uuid.UUID{}
	lr, err := r.svc.Get(server.Context())
	if err != nil {
		return status.Errorf(codes.Internal,
			fmt.Sprintf("Unable to retreive report: %v", err))
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
	err = server.Send(&reportpb.GetResponse{Report: frep})
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
					fmt.Sprintf("Unable to retreive report: %v", err))
			}
			frep, err = r.filterReport(rol, tID, lr)
			if err != nil {
				return status.Errorf(codes.Internal,
					fmt.Sprintf("Unable to filter report: %v", err))
			}
			err = server.Send(&reportpb.GetResponse{Report: frep})
			if err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

func NewReportController(svc report_service.Serv, reportClient *report_client.Client, client *policy_client.Client) *reportController {
	return &reportController{
		svc:          svc,
		reportClient: reportClient,
		policyClient: client,
	}
}
