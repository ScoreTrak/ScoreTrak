package handler

import (
	"context"

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
	checkv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v1"
	competitionv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/competition/v1"
	hostv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host/v1"
	host_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v1"
	propertyv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/property/v1"
	reportv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/report/v1"
	roundv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/round/v1"
	servicev1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service/v1"
	service_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service_group/v1"
	teamv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v1"
	userv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CompetitionController struct {
	svc competitionservice.Serv
	competitionv1.UnimplementedCompetitionServiceServer
}

func (c CompetitionController) LoadCompetition(ctx context.Context, request *competitionv1.LoadCompetitionRequest) (*competitionv1.LoadCompetitionResponse, error) {
	hstGrps := make([]*hostgroup.HostGroup, 0, len(request.GetCompetition().HostGroups))
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
	servGrps := make([]*servicegroup.ServiceGroup, 0, len(request.GetCompetition().ServiceGroups))
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
	return &competitionv1.LoadCompetitionResponse{}, err
}

func (c CompetitionController) FetchCoreCompetition(ctx context.Context, _ *competitionv1.FetchCoreCompetitionRequest) (*competitionv1.FetchCoreCompetitionResponse, error) {
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
	return &competitionv1.FetchCoreCompetitionResponse{Competition: comppb}, nil
}

func (c CompetitionController) FetchEntireCompetition(ctx context.Context, _ *competitionv1.FetchEntireCompetitionRequest) (*competitionv1.FetchEntireCompetitionResponse, error) {
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
	return &competitionv1.FetchEntireCompetitionResponse{Competition: comppb}, nil
}

func (c CompetitionController) ResetScores(ctx context.Context, _ *competitionv1.ResetScoresRequest) (*competitionv1.ResetScoresResponse, error) {
	if err := c.svc.ResetScores(ctx); err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	return &competitionv1.ResetScoresResponse{}, nil
}

func (c CompetitionController) DeleteCompetition(ctx context.Context, _ *competitionv1.DeleteCompetitionRequest) (*competitionv1.DeleteCompetitionResponse, error) {
	err := c.svc.DeleteCompetition(ctx)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"Internal Error when fetching competition: %v", err,
		)
	}
	return &competitionv1.DeleteCompetitionResponse{}, nil
}

func NewCompetitionController(svc competitionservice.Serv) *CompetitionController {
	return &CompetitionController{svc: svc}
}

func ConvertCompetitionToCompetitionPB(comp *competition.Competition) (*competitionv1.Competition, error) {
	hstGrps := make([]*host_groupv1.HostGroup, 0, len(comp.HostGroups))
	for i := range comp.HostGroups {
		hstGrps = append(hstGrps, ConvertHostGroupToHostGroupPb(comp.HostGroups[i]))
	}
	hsts := make([]*hostv1.Host, 0, len(comp.Hosts))
	for i := range comp.Hosts {
		hsts = append(hsts, ConvertHostToHostPb(comp.Hosts[i]))
	}
	tms := make([]*teamv1.Team, 0, len(comp.Teams))
	for i := range comp.Teams {
		tms = append(tms, ConvertTeamToTeamPb(comp.Teams[i]))
	}
	svcs := make([]*servicev1.Service, 0, len(comp.Services))
	for i := range comp.Services {
		svcs = append(svcs, ConvertServiceToServicePb(comp.Services[i]))
	}
	servGrps := make([]*service_groupv1.ServiceGroup, 0, len(comp.ServiceGroups))
	for i := range comp.ServiceGroups {
		servGrps = append(servGrps, ConvertServiceGroupToServiceGroupPb(comp.ServiceGroups[i]))
	}
	rnds := make([]*roundv1.Round, 0, len(comp.Rounds))
	for i := range comp.Rounds {
		rnds = append(rnds, ConvertRoundToRoundPb(comp.Rounds[i]))
	}
	props := make([]*propertyv1.Property, 0, len(comp.Properties))
	for i := range comp.Properties {
		props = append(props, ConvertPropertyToPropertyPb(comp.Properties[i]))
	}
	chcks := make([]*checkv1.Check, 0, len(comp.Checks))
	for i := range comp.Checks {
		chcks = append(chcks, ConvertCheckToCheckPb(comp.Checks[i]))
	}
	usrs := make([]*userv1.User, 0, len(comp.Users))
	for i := range comp.Users {
		usrs = append(usrs, ConvertUserToUserPb(comp.Users[i]))
	}

	var rprt *reportv1.Report
	if comp.Report != nil {
		rprt = &reportv1.Report{
			Cache:     comp.Report.Cache,
			UpdatedAt: timestamppb.New(comp.Report.UpdatedAt),
		}
	}

	return &competitionv1.Competition{
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
