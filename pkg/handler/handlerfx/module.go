package handlerfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/handler"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/serviceservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teamservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	healthv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/grpc/health/v1"
	authv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/auth/v1"
	checkv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v1"
	competitionv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/competition/v1"
	configv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/config/v1"
	hostv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host/v1"
	host_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v1"
	policyv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/policy/v1"
	propertyv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/property/v1"
	reportv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/report/v1"
	roundv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/round/v1"
	servicev1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service/v1"
	service_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service_group/v1"
	teamv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v1"
	userv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v1"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

var Module = fx.Options(
	// Add Service Pattern Implementations
	fx.Provide(
		checkservice.NewCheckServ,
		hostgroupservice.NewHostGroupServ,
		hostservice.NewHostServ,
		propertyservice.NewPropertyServ,
		configservice.NewConfigServ,
		configservice.NewStaticConfigServ,
		competitionservice.NewCompetitionServ,
		policyservice.NewPolicyServ,
		reportservice.NewReportServ,
		userservice.NewUserServ,
		teamservice.NewTeamServ,
		servicegroupservice.NewServiceGroupServ,
		roundservice.NewRoundServ,
		serviceservice.NewServiceServ,
	),

	// Add Handlers/grpc server
	fx.Provide(
		fx.Annotate(handler.NewHealthController, fx.As(new(healthv1.HealthServer))),
		fx.Annotate(handler.NewPolicyController, fx.As(new(policyv1.PolicyServiceServer))),
		fx.Annotate(handler.NewAuthController, fx.As(new(authv1.AuthServiceServer))),
		fx.Annotate(handler.NewUserController, fx.As(new(userv1.UserServiceServer))),
		fx.Annotate(handler.NewTeamController, fx.As(new(teamv1.TeamServiceServer))),
		fx.Annotate(handler.NewServiceController, fx.As(new(servicev1.ServiceServiceServer))),
		fx.Annotate(handler.NewServiceGroupController, fx.As(new(service_groupv1.ServiceGroupServiceServer))),
		fx.Annotate(handler.NewRoundController, fx.As(new(roundv1.RoundServiceServer))),
		fx.Annotate(handler.NewReportController, fx.As(new(reportv1.ReportServiceServer))),
		fx.Annotate(handler.NewPropertyController, fx.As(new(propertyv1.PropertyServiceServer))),
		fx.Annotate(handler.NewHostGroupController, fx.As(new(host_groupv1.HostGroupServiceServer))),
		fx.Annotate(handler.NewHostController, fx.As(new(hostv1.HostServiceServer))),
		fx.Annotate(handler.NewStaticConfigController, fx.As(new(configv1.StaticConfigServiceServer))),
		fx.Annotate(handler.NewConfigController, fx.As(new(configv1.DynamicConfigServiceServer))),
		fx.Annotate(handler.NewCompetitionController, fx.As(new(competitionv1.CompetitionServiceServer))),
		fx.Annotate(handler.NewCheckController, fx.As(new(checkv1.CheckServiceServer))),
	),

	fx.Provide(func(server *grpc.Server) grpc.ServiceRegistrar {
		return server
	}),
	// Add them to the grpc server object
	fx.Invoke(
		checkv1.RegisterCheckServiceServer,
		competitionv1.RegisterCompetitionServiceServer,
		configv1.RegisterStaticConfigServiceServer,
		configv1.RegisterDynamicConfigServiceServer,
		hostv1.RegisterHostServiceServer,
		host_groupv1.RegisterHostGroupServiceServer,
		propertyv1.RegisterPropertyServiceServer,
		reportv1.RegisterReportServiceServer,
		roundv1.RegisterRoundServiceServer,
		servicev1.RegisterServiceServiceServer,
		service_groupv1.RegisterServiceGroupServiceServer,
		teamv1.RegisterTeamServiceServer,
		userv1.RegisterUserServiceServer,
		authv1.RegisterAuthServiceServer,
		policyv1.RegisterPolicyServiceServer,
		healthv1.RegisterHealthServer,
	),
)
