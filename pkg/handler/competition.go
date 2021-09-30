package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competition_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	checkpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/check/v1"
	competitionpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/competition/v1"
	hostpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/host/v1"
	host_grouppb "github.com/ScoreTrak/ScoreTrak/pkg/proto/host_group/v1"
	propertypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/property/v1"
	reportpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/report/v1"
	roundpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/round/v1"
	servicepb "github.com/ScoreTrak/ScoreTrak/pkg/proto/service/v1"
	service_grouppb "github.com/ScoreTrak/ScoreTrak/pkg/proto/service_group/v1"
	teampb "github.com/ScoreTrak/ScoreTrak/pkg/proto/team/v1"
	userpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/user/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompetitionController struct {
	svc competition_service.Serv
	competitionpb.UnimplementedCompetitionServiceServer
}

func (c CompetitionController) LoadCompetition(ctx context.Context, request *competitionpb.LoadCompetitionRequest) (*competitionpb.LoadCompetitionResponse, error) {
	hstGrps := make([]*host_group.HostGroup, 0, len(request.GetCompetition().HostGroups))
	for i := range request.GetCompetition().HostGroups {
		hstGrp, err := ConvertHostGroupPBtoHostGroup(true, request.GetCompetition().HostGroups[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" host groups, details: %v", err)
		}
		hstGrps = append(hstGrps, hstGrp)
	}
	hsts := make([]*host.Host, 0, len(request.GetCompetition().Hosts))
	for i := range request.GetCompetition().Hosts {
		hst, err := ConvertHostPBtoHost(true, request.GetCompetition().Hosts[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" host, details: %v", err)
		}
		hsts = append(hsts, hst)
	}
	tms := make([]*team.Team, 0, len(request.GetCompetition().Teams))
	for i := range request.GetCompetition().Teams {
		tm, err := ConvertTeamPBtoTeam(true, request.GetCompetition().Teams[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" team, details: %v", err)
		}
		tms = append(tms, tm)
	}
	svcs := make([]*service2.Service, 0, len(request.GetCompetition().Services))
	for i := range request.GetCompetition().Services {
		svc, err := ConvertServicePBtoService(true, request.GetCompetition().Services[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" check_service groups, details: %v", err)
		}
		svcs = append(svcs, svc)
	}
	servGrps := make([]*service_group.ServiceGroup, 0, len(request.GetCompetition().ServiceGroups))
	for i := range request.GetCompetition().ServiceGroups {
		servGrp, err := ConvertServiceGroupPBtoServiceGroup(true, request.GetCompetition().ServiceGroups[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" service groups, details: %v", err)
		}
		servGrps = append(servGrps, servGrp)
	}
	rnds := make([]*round.Round, 0, len(request.GetCompetition().Rounds))
	for i := range request.GetCompetition().Rounds {
		rnd, err := ConvertRoundPBtoRound(true, request.GetCompetition().Rounds[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" round, details: %v", err)
		}
		rnds = append(rnds, rnd)
	}
	props := make([]*property.Property, 0, len(request.GetCompetition().Properties))
	for i := range request.GetCompetition().Properties {
		prop, err := ConvertPropertyPBtoProperty(request.GetCompetition().Properties[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" property, details: %v", err)
		}
		props = append(props, prop)
	}
	chcks := make([]*check.Check, 0, len(request.GetCompetition().Checks))
	for i := range request.GetCompetition().Checks {
		chck, err := ConvertCheckPBtoCheck(request.GetCompetition().Checks[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" check, details: %v", err)
		}
		chcks = append(chcks, chck)
	}

	users := make([]*user.User, 0, len(request.GetCompetition().Users))
	for i := range request.GetCompetition().Users {
		usr, err := ConvertUserPBtoUser(true, request.GetCompetition().Users[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" check, details: %v", err)
		}
		users = append(users, usr)
	}

	var cache string
	if request.Competition.Report != nil {
		cache = request.Competition.Report.Cache
	}

	err := c.svc.LoadCompetition(ctx, &competition.Competition{
		Config: ConvertDynamicConfigPBToDynamicConfig(request.GetCompetition().DynamicConfig),
		Report: &report.Report{
			Cache: cache,
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
		Policy:        ConvertPolicyPBToPolicy(request.GetCompetition().Policy),
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
	if err := c.svc.ResetScores(ctx); err != nil {
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

func NewCompetitionController(svc competition_service.Serv) *CompetitionController {
	return &CompetitionController{svc: svc}
}

func ConvertCompetitionToCompetitionPB(comp *competition.Competition) (*competitionpb.Competition, error) {
	hstGrps := make([]*host_grouppb.HostGroup, 0, len(comp.HostGroups))
	for i := range comp.HostGroups {
		hstGrps = append(hstGrps, ConvertHostGroupToHostGroupPb(comp.HostGroups[i]))
	}
	hsts := make([]*hostpb.Host, 0, len(comp.Hosts))
	for i := range comp.Hosts {
		hsts = append(hsts, ConvertHostToHostPb(comp.Hosts[i]))
	}
	tms := make([]*teampb.Team, 0, len(comp.Teams))
	for i := range comp.Teams {
		tms = append(tms, ConvertTeamToTeamPb(comp.Teams[i]))
	}
	svcs := make([]*servicepb.Service, 0, len(comp.Services))
	for i := range comp.Services {
		svcs = append(svcs, ConvertServiceToServicePb(comp.Services[i]))
	}
	servGrps := make([]*service_grouppb.ServiceGroup, 0, len(comp.ServiceGroups))
	for i := range comp.ServiceGroups {
		servGrps = append(servGrps, ConvertServiceGroupToServiceGroupPb(comp.ServiceGroups[i]))
	}
	rnds := make([]*roundpb.Round, 0, len(comp.Rounds))
	for i := range comp.Rounds {
		rnds = append(rnds, ConvertRoundToRoundPb(comp.Rounds[i]))
	}
	props := make([]*propertypb.Property, 0, len(comp.Properties))
	for i := range comp.Properties {
		props = append(props, ConvertPropertyToPropertyPb(comp.Properties[i]))
	}
	chcks := make([]*checkpb.Check, 0, len(comp.Checks))
	for i := range comp.Checks {
		chcks = append(chcks, ConvertCheckToCheckPb(comp.Checks[i]))
	}
	usrs := make([]*userpb.User, 0, len(comp.Users))
	for i := range comp.Users {
		usrs = append(usrs, ConvertUserToUserPb(comp.Users[i]))
	}

	var rprt *reportpb.Report
	if comp.Report != nil {
		rprt = &reportpb.Report{
			Cache:     comp.Report.Cache,
			UpdatedAt: timestamppb.New(comp.Report.UpdatedAt),
		}
	}

	return &competitionpb.Competition{
		DynamicConfig: ConvertDynamicConfigToDynamicConfigPB(comp.Config),
		Report:        rprt,
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
