package handler

import (
	"encoding/json"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_client"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type reportController struct {
	svc          report_service.Serv
	reportClient *report_client.Client
	policyClient *policy_client.Client
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
			if !simpleReport.Teams[t].Enabled || simpleReport.Teams[t].Hidden {
				delete(simpleReport.Teams, t)
				continue
			}
			for h := range simpleReport.Teams[t].Hosts {
				if !simpleReport.Teams[t].Hosts[h].Enabled || (simpleReport.Teams[t].Hosts[h].HostGroup != nil && !simpleReport.Teams[t].Hosts[h].HostGroup.Enabled) {
					delete(simpleReport.Teams[t].Hosts, h)
					continue
				}
				for s := range simpleReport.Teams[t].Hosts[h].Services {
					if !simpleReport.Teams[t].Hosts[h].Services[s].Enabled || (!simpleReport.Teams[t].Hosts[h].Services[s].SimpleServiceGroup.Enabled) {
						delete(simpleReport.Teams[t].Hosts[h].Services, s)
						continue
					}
				}
			}
		}
	}
	if rol != role.Black {
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
						simpleReport.Teams[t].Hosts[h].Services[s].Err = ""
						simpleReport.Teams[t].Hosts[h].Services[s].Log = ""
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
	rol := role.Anonymous
	tID := uuid.UUID{}
	lr, err := r.svc.Get(server.Context())
	if err != nil {
		return status.Errorf(codes.Internal,
			fmt.Sprintf("Unable to retreive report: %v", err))
	}

	if val, ok := server.Context().Value("claims").(*auth.UserClaims); ok && val != nil {
		rol = val.Role
		tID = uuid.FromStringOrNil(val.TeamID)
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
	authTimer := time.NewTimer(time.Second * 30)
	defer r.reportClient.Unsubscribe(uid)

	for {
		select {
		case <-ch:
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
		case <-authTimer.C:
			if rol == role.Anonymous && !r.policyClient.GetAllowUnauthenticatedUsers() {
				return status.Error(codes.PermissionDenied, "You must login in order to access this resource")
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
