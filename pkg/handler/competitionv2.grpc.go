package handler

import (
	"context"
	checkv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v2"
	competitionv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/competition/v2"
	hostv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host/v2"
	host_groupv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v2"
	propertyv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/property/v2"
	reportv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/report/v2"
	roundv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/round/v2"
	servicev2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service/v2"
	service_groupv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service_group/v2"
	teamv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v2"
	userv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v2"

	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/host"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/property"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
	service2 "github.com/ScoreTrak/ScoreTrak/pkg/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompetitionV2Controller struct {
	svc competitionservice.Serv
	competitionv2.UnimplementedCompetitionServiceServer
}

func (c CompetitionV2Controller) LoadCompetition(ctx context.Context, request *competitionv2.CompetitionServiceLoadCompetitionRequest) (*competitionv2.CompetitionServiceLoadCompetitionResponse, error) {
	hstGrps := make([]*hostgroup.HostGroup, 0, len(request.GetCompetition().HostGroups))
	for i := range request.GetCompetition().HostGroups {
		hstGrp, err := ConvertHostGroupV2PBtoHostGroup(true, request.GetCompetition().HostGroups[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" host groups, details: %v", err)
		}
		hstGrps = append(hstGrps, hstGrp)
	}
	hsts := make([]*host.Host, 0, len(request.GetCompetition().Hosts))
	for i := range request.GetCompetition().Hosts {
		hst, err := ConvertHostV2PBtoHost(true, request.GetCompetition().Hosts[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" host, details: %v", err)
		}
		hsts = append(hsts, hst)
	}
	tms := make([]*team.Team, 0, len(request.GetCompetition().Teams))
	for i := range request.GetCompetition().Teams {
		tm, err := ConvertTeamV2PBtoTeam(true, request.GetCompetition().Teams[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" team, details: %v", err)
		}
		tms = append(tms, tm)
	}
	svcs := make([]*service2.Service, 0, len(request.GetCompetition().Services))
	for i := range request.GetCompetition().Services {
		svc, err := ConvertServiceV2PBtoService(true, request.GetCompetition().Services[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" check_service groups, details: %v", err)
		}
		svcs = append(svcs, svc)
	}
	servGrps := make([]*servicegroup.ServiceGroup, 0, len(request.GetCompetition().ServiceGroups))
	for i := range request.GetCompetition().ServiceGroups {
		servGrp, err := ConvertServiceGroupV2PBtoServiceGroup(true, request.GetCompetition().ServiceGroups[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" service groups, details: %v", err)
		}
		servGrps = append(servGrps, servGrp)
	}
	rnds := make([]*round.Round, 0, len(request.GetCompetition().Rounds))
	for i := range request.GetCompetition().Rounds {
		rnd, err := ConvertRoundV2PBtoRound(true, request.GetCompetition().Rounds[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" round, details: %v", err)
		}
		rnds = append(rnds, rnd)
	}
	props := make([]*property.Property, 0, len(request.GetCompetition().Properties))
	for i := range request.GetCompetition().Properties {
		prop, err := ConvertPropertyV2PBtoProperty(request.GetCompetition().Properties[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" property, details: %v", err)
		}
		props = append(props, prop)
	}
	chcks := make([]*check.Check, 0, len(request.GetCompetition().Checks))
	for i := range request.GetCompetition().Checks {
		chck, err := ConvertCheckV2PBtoCheck(request.GetCompetition().Checks[i])
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, unableToParse+" check, details: %v", err)
		}
		chcks = append(chcks, chck)
	}

	users := make([]*user.User, 0, len(request.GetCompetition().Users))
	for i := range request.GetCompetition().Users {
		usr, err := ConvertUserV2PBtoUser(true, request.GetCompetition().Users[i])
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
		Config: ConvertDynamicConfigV2PBToDynamicConfig(request.GetCompetition().DynamicConfig),
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
		Policy:        ConvertPolicyV2PBToPolicy(request.GetCompetition().Policy),
	})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when loading competition: %v", err,
		)
	}
	return &competitionv2.CompetitionServiceLoadCompetitionResponse{}, err
}

func (c CompetitionV2Controller) FetchCoreCompetition(ctx context.Context, _ *competitionv2.CompetitionServiceFetchCoreCompetitionRequest) (*competitionv2.CompetitionServiceFetchCoreCompetitionResponse, error) {
	comp, err := c.svc.FetchCoreCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	comppb, err := ConvertCompetitionToCompetitionV2PB(comp)
	if err != nil {
		return nil, err
	}
	return &competitionv2.CompetitionServiceFetchCoreCompetitionResponse{Competition: comppb}, nil
}

func (c CompetitionV2Controller) FetchEntireCompetition(ctx context.Context, _ *competitionv2.CompetitionServiceFetchEntireCompetitionRequest) (*competitionv2.CompetitionServiceFetchEntireCompetitionResponse, error) {
	comp, err := c.svc.FetchEntireCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	comppb, err := ConvertCompetitionToCompetitionV2PB(comp)
	if err != nil {
		return nil, err
	}
	return &competitionv2.CompetitionServiceFetchEntireCompetitionResponse{Competition: comppb}, nil
}

func (c CompetitionV2Controller) ResetScores(ctx context.Context, _ *competitionv2.CompetitionServiceResetScoresRequest) (*competitionv2.CompetitionServiceResetScoresResponse, error) {
	if err := c.svc.ResetScores(ctx); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	return &competitionv2.CompetitionServiceResetScoresResponse{}, nil
}

func (c CompetitionV2Controller) DeleteCompetition(ctx context.Context, _ *competitionv2.CompetitionServiceDeleteCompetitionRequest) (*competitionv2.CompetitionServiceDeleteCompetitionResponse, error) {
	err := c.svc.DeleteCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	return &competitionv2.CompetitionServiceDeleteCompetitionResponse{}, nil
}

func NewCompetitionV2Controller(svc competitionservice.Serv) *CompetitionV2Controller {
	return &CompetitionV2Controller{svc: svc}
}

func ConvertCompetitionToCompetitionV2PB(comp *competition.Competition) (*competitionv2.Competition, error) {
	hstGrps := make([]*host_groupv2.HostGroup, 0, len(comp.HostGroups))
	for i := range comp.HostGroups {
		hstGrps = append(hstGrps, ConvertHostGroupToHostGroupV2Pb(comp.HostGroups[i]))
	}
	hsts := make([]*hostv2.Host, 0, len(comp.Hosts))
	for i := range comp.Hosts {
		hsts = append(hsts, ConvertHostToHostV2Pb(comp.Hosts[i]))
	}
	tms := make([]*teamv2.Team, 0, len(comp.Teams))
	for i := range comp.Teams {
		tms = append(tms, ConvertTeamToTeamV2Pb(comp.Teams[i]))
	}
	svcs := make([]*servicev2.Service, 0, len(comp.Services))
	for i := range comp.Services {
		svcs = append(svcs, ConvertServiceToServiceV2Pb(comp.Services[i]))
	}
	servGrps := make([]*service_groupv2.ServiceGroup, 0, len(comp.ServiceGroups))
	for i := range comp.ServiceGroups {
		servGrps = append(servGrps, ConvertServiceGroupToServiceGroupV2Pb(comp.ServiceGroups[i]))
	}
	rnds := make([]*roundv2.Round, 0, len(comp.Rounds))
	for i := range comp.Rounds {
		rnds = append(rnds, ConvertRoundToRoundV2Pb(comp.Rounds[i]))
	}
	props := make([]*propertyv2.Property, 0, len(comp.Properties))
	for i := range comp.Properties {
		props = append(props, ConvertPropertyToPropertyV2Pb(comp.Properties[i]))
	}
	chcks := make([]*checkv2.Check, 0, len(comp.Checks))
	for i := range comp.Checks {
		chcks = append(chcks, ConvertCheckToCheckV2Pb(comp.Checks[i]))
	}
	usrs := make([]*userv2.User, 0, len(comp.Users))
	for i := range comp.Users {
		usrs = append(usrs, ConvertUserToUserV2Pb(comp.Users[i]))
	}

	var rprt *reportv2.Report
	if comp.Report != nil {
		rprt = &reportv2.Report{
			Cache:     comp.Report.Cache,
			UpdatedAt: timestamppb.New(comp.Report.UpdatedAt),
		}
	}

	return &competitionv2.Competition{
		DynamicConfig: ConvertDynamicConfigToDynamicConfigV2PB(comp.Config),
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
		Policy:        ConvertPolicyToPolicyV2PB(comp.Policy),
	}, nil
}
