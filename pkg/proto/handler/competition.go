package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_grouppb"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertypb"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundpb"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicepb"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_grouppb"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teampb"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userpb"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompetitionController struct {
	svc service.Serv
}

func (c CompetitionController) LoadCompetition(ctx context.Context, request *competitionpb.LoadCompetitionRequest) (*competitionpb.LoadCompetitionResponse, error) {
	var hstGrps []*host_group.HostGroup
	for i := range request.Competition.HostGroups {
		hstGrp, err := ConvertHostGroupPBtoHostGroup(true, request.Competition.HostGroups[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse host groups, details: %v", err)
		}
		hstGrps = append(hstGrps, hstGrp)
	}
	var hsts []*host.Host
	for i := range request.Competition.Hosts {
		hst, err := ConvertHostPBtoHost(true, request.Competition.Hosts[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse host, details: %v", err)
		}
		hsts = append(hsts, hst)
	}
	var tms []*team.Team
	for i := range request.Competition.Teams {
		tm, err := ConvertTeamPBtoTeam(true, request.Competition.Teams[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse team, details: %v", err)
		}
		tms = append(tms, tm)
	}
	var svcs []*service2.Service
	for i := range request.Competition.Services {
		svc, err := ConvertServicePBtoService(true, request.Competition.Services[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse service groups, details: %v", err)
		}
		svcs = append(svcs, svc)
	}
	var servGrps []*service_group.ServiceGroup
	for i := range request.Competition.ServiceGroups {
		servGrp, err := ConvertServiceGroupPBtoServiceGroup(true, request.Competition.ServiceGroups[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse host groups, details: %v", err)
		}
		servGrps = append(servGrps, servGrp)
	}
	var rnds []*round.Round
	for i := range request.Competition.Rounds {
		rnd, err := ConvertRoundPBtoRound(true, request.Competition.Rounds[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse round, details: %v", err)
		}
		rnds = append(rnds, rnd)
	}
	var props []*property.Property
	for i := range request.Competition.Properties {
		prop, err := ConvertPropertyPBtoProperty(request.Competition.Properties[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse property, details: %v", err)
		}
		props = append(props, prop)
	}
	var chcks []*check.Check
	for i := range request.Competition.Checks {
		chck, err := ConvertCheckPBtoCheck(request.Competition.Checks[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse check, details: %v", err)
		}
		chcks = append(chcks, chck)
	}

	var users []*user.User
	for i := range request.Competition.Users {
		usr, err := ConvertUserPBtoUser(true, request.Competition.Users[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse check, details: %v", err)
		}
		users = append(users, usr)
	}

	err := c.svc.LoadCompetition(ctx, &competition.Competition{
		Config: ConvertDynamicConfigPBToDynamicConfig(request.Competition.DynamicConfig),
		Report: &report.Report{
			Cache: request.Competition.Report.Cache,
		},
		HostGroups:    hstGrps,
		Hosts:         hsts,
		Teams:         tms,
		Services:      svcs,
		ServiceGroups: servGrps,
		Rounds:        rnds,
		Properties:    props,
		Checks:        chcks,
		Users:         users,
		Policy:        ConvertPolicyPBToPolicy(request.Competition.Policy),
	})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when loading competition: %v", err,
		)
	}
	return &competitionpb.LoadCompetitionResponse{}, err
}

func (c CompetitionController) FetchCoreCompetition(ctx context.Context, request *competitionpb.FetchCoreCompetitionRequest) (*competitionpb.FetchCoreCompetitionResponse, error) {
	comp, err := c.svc.FetchCoreCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	comppb, err := ConvertCompetitionToCompetitionPB(comp)
	if err != nil {
		return nil, err
	}
	return &competitionpb.FetchCoreCompetitionResponse{Competition: comppb}, nil
}

func (c CompetitionController) FetchEntireCompetition(ctx context.Context, request *competitionpb.FetchEntireCompetitionRequest) (*competitionpb.FetchEntireCompetitionResponse, error) {
	comp, err := c.svc.FetchEntireCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	comppb, err := ConvertCompetitionToCompetitionPB(comp)
	if err != nil {
		return nil, err
	}
	return &competitionpb.FetchEntireCompetitionResponse{Competition: comppb}, nil
}

func (c CompetitionController) ResetScores(ctx context.Context, request *competitionpb.ResetScoresRequest) (*competitionpb.ResetScoresResponse, error) {
	err := c.svc.ResetScores(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	return &competitionpb.ResetScoresResponse{}, nil
}

func (c CompetitionController) DeleteCompetition(ctx context.Context, request *competitionpb.DeleteCompetitionRequest) (*competitionpb.DeleteCompetitionResponse, error) {
	err := c.svc.DeleteCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	return &competitionpb.DeleteCompetitionResponse{}, nil
}

func NewCompetitionController(svc service.Serv) *CompetitionController {
	return &CompetitionController{svc}
}

func ConvertCompetitionToCompetitionPB(comp *competition.Competition) (*competitionpb.Competition, error) {

	var hstGrps []*host_grouppb.HostGroup
	for i := range comp.HostGroups {
		hstGrps = append(hstGrps, ConvertHostGroupToHostGroupPb(comp.HostGroups[i]))
	}
	var hsts []*hostpb.Host
	for i := range comp.Hosts {
		hsts = append(hsts, ConvertHostToHostPb(comp.Hosts[i]))
	}
	var tms []*teampb.Team
	for i := range comp.Teams {
		tms = append(tms, ConvertTeamToTeamPb(comp.Teams[i]))
	}
	var svcs []*servicepb.Service
	for i := range comp.Services {
		svcs = append(svcs, ConvertServiceToServicePb(comp.Services[i]))
	}
	var servGrps []*service_grouppb.ServiceGroup
	for i := range comp.ServiceGroups {
		servGrps = append(servGrps, ConvertServiceGroupToServiceGroupPb(comp.ServiceGroups[i]))
	}
	var rnds []*roundpb.Round
	for i := range comp.Rounds {
		rnd, err := ConvertRoundToRoundPb(comp.Rounds[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "Unable to parse round, details: %v", err)
		}
		rnds = append(rnds, rnd)
	}
	var props []*propertypb.Property
	for i := range comp.Properties {
		props = append(props, ConvertPropertyToPropertyPb(comp.Properties[i]))
	}
	var chcks []*checkpb.Check
	for i := range comp.Checks {
		chcks = append(chcks, ConvertCheckToCheckPb(comp.Checks[i]))
	}
	var usrs []*userpb.User
	for i := range comp.Users {
		usrs = append(usrs, ConvertUserToUserPb(comp.Users[i]))
	}

	uat, err := ptypes.TimestampProto(comp.Report.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Unable convert time.date to timestamp(Ideally, this should not happen, perhaps this is a bug): %v", err),
		)
	}

	return &competitionpb.Competition{
		DynamicConfig: ConvertDynamicConfigToDynamicConfigPB(comp.Config),
		Report: &reportpb.Report{
			Cache:     comp.Report.Cache,
			UpdatedAt: uat,
		},
		HostGroups:    hstGrps,
		Hosts:         hsts,
		Teams:         tms,
		Services:      svcs,
		ServiceGroups: servGrps,
		Rounds:        rnds,
		Properties:    props,
		Checks:        chcks,
		Users:         usrs,
		Policy:        ConvertPolicyToPolicyPB(comp.Policy),
	}, nil
}