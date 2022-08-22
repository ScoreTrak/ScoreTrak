package handler

import (
	"encoding/json"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportclient"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	reportv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/report/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ReportController struct {
	svc          reportservice.Serv
	reportClient *reportclient.Client
	policyClient *policyclient.Client
	reportv1.UnimplementedReportServiceServer
}

func removeDisabledAndHidden(simpleReport *report.SimpleReport) {
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

func calculateTotalPoints(simpleReport *report.SimpleReport) {
	for t := range simpleReport.Teams {
		for h := range simpleReport.Teams[t].Hosts {
			for s := range simpleReport.Teams[t].Hosts[h].Services {
				simpleReport.Teams[t].TotalPoints += simpleReport.Teams[t].Hosts[h].Services[s].Points + simpleReport.Teams[t].Hosts[h].Services[s].PointsBoost
			}
		}
	}
}

func filterBlueTeams(simpleReport *report.SimpleReport, tID uuid.UUID, p *policy.Policy) {
	for teamIdx := range simpleReport.Teams {
		for hostIdx := range simpleReport.Teams[teamIdx].Hosts {
			for serviceIdx := range simpleReport.Teams[teamIdx].Hosts[hostIdx].Services {
				propFilterHide := map[string]*report.SimpleProperty{}
				for key, val := range simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Properties {
					if val.Status != "Hide" {
						propFilterHide[key] = val
					}
				}
				simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Properties = propFilterHide
				if teamIdx != tID {
					if simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Check != nil {
						simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Check.Err = ""
						simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Check.Log = ""
					}
					prop := map[string]*report.SimpleProperty{}
					if *p.ShowAddresses {
						if val, ok := simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Properties["Port"]; ok {
							prop["Port"] = val
						}
					}
					simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Properties = prop
					if !*p.ShowPoints {
						simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].Points = 0
						simpleReport.Teams[teamIdx].Hosts[hostIdx].Services[serviceIdx].PointsBoost = 0
						simpleReport.Teams[teamIdx].TotalPoints = 0
					}
				}
			}
			if teamIdx != tID {
				if !*p.ShowAddresses {
					simpleReport.Teams[teamIdx].Hosts[hostIdx].Address = ""
				}
			}
		}
	}
}

func (r *ReportController) filterReport(rol string, tID uuid.UUID, lr *report.Report) (*reportv1.Report, error) {
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
	return &reportv1.Report{
		Cache:     string(ret),
		UpdatedAt: timestamppb.New(lr.UpdatedAt),
	}, nil
}

func (r *ReportController) Get(_ *reportv1.GetRequest, server reportv1.ReportService_GetServer) error {
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
	err = server.Send(&reportv1.GetResponse{Report: frep})
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
			err = server.Send(&reportv1.GetResponse{Report: frep})
			if err != nil {
				return err
			}
		case <-server.Context().Done():
			return nil
		}
	}
}

func NewReportController(svc reportservice.Serv, reportClient *reportclient.Client, client *policyclient.Client) *ReportController {
	return &ReportController{
		svc:          svc,
		reportClient: reportClient,
		policyClient: client,
	}
}
