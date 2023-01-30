package handlerfx

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/handler"
	healthv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/grpc/health/v1"
	authv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/auth/v1"
	authv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/auth/v2"
	checkv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v1"
	checkv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/check/v2"
	competitionv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/competition/v1"
	competitionv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/competition/v2"
	configv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/config/v1"
	configv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/config/v2"
	hostv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host/v1"
	hostv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host/v2"
	host_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v1"
	host_groupv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/host_group/v2"
	policyv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/policy/v1"
	policyv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/policy/v2"
	propertyv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/property/v1"
	propertyv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/property/v2"
	reportv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/report/v1"
	reportv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/report/v2"
	roundv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/round/v1"
	roundv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/round/v2"
	servicev1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service/v1"
	servicev2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service/v2"
	service_groupv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service_group/v1"
	service_groupv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/service_group/v2"
	teamv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v1"
	teamv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/team/v2"
	userv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v1"
	userv2 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v2"
	"go.uber.org/fx"
)

//var ConnectModule = fx.Options(
//	// Create connect handlers
//	fx.Provide(
//		fx.Annotate(handler.NewAuthV1ConnectServer, fx.As(new(authv1connect.AuthServiceHandler))),
//		fx.Annotate(handler.NewAuthV2ConnectServer, fx.As(new(authv2connect.AuthServiceHandler))),
//	),
//
//	// Register them to the connect server
//	fx.Provide(
//		authv1connect.NewAuthServiceHandler,
//		authv2connect.NewAuthServiceHandler,
//	),
//)

var GrpcModule = fx.Options(
	// Create controller/handlers
	fx.Provide(
		fx.Annotate(handler.NewHealthV1Controller, fx.As(new(healthv1.HealthServer))),
		fx.Annotate(handler.NewPolicyV1Controller, fx.As(new(policyv1.PolicyServiceServer))),
		fx.Annotate(handler.NewPolicyV2Controller, fx.As(new(policyv2.PolicyServiceServer))),
		fx.Annotate(handler.NewAuthV1Controller, fx.As(new(authv1.AuthServiceServer))),
		fx.Annotate(handler.NewAuthV2Controller, fx.As(new(authv2.AuthServiceServer))),
		fx.Annotate(handler.NewUserV1Controller, fx.As(new(userv1.UserServiceServer))),
		fx.Annotate(handler.NewUserV2Controller, fx.As(new(userv2.UserServiceServer))),
		fx.Annotate(handler.NewTeamV1Controller, fx.As(new(teamv1.TeamServiceServer))),
		fx.Annotate(handler.NewTeamV2Controller, fx.As(new(teamv2.TeamServiceServer))),
		fx.Annotate(handler.NewServiceV1Controller, fx.As(new(servicev1.ServiceServiceServer))),
		fx.Annotate(handler.NewServiceV2Controller, fx.As(new(servicev2.ServiceServiceServer))),
		fx.Annotate(handler.NewServiceGroupV1Controller, fx.As(new(service_groupv1.ServiceGroupServiceServer))),
		fx.Annotate(handler.NewServiceGroupV2Controller, fx.As(new(service_groupv2.ServiceGroupServiceServer))),
		fx.Annotate(handler.NewRoundV1Controller, fx.As(new(roundv1.RoundServiceServer))),
		fx.Annotate(handler.NewRoundV2Controller, fx.As(new(roundv2.RoundServiceServer))),
		fx.Annotate(handler.NewReportV1Controller, fx.As(new(reportv1.ReportServiceServer))),
		fx.Annotate(handler.NewReportV2Controller, fx.As(new(reportv2.ReportServiceServer))),
		fx.Annotate(handler.NewPropertyV1Controller, fx.As(new(propertyv1.PropertyServiceServer))),
		fx.Annotate(handler.NewPropertyV2Controller, fx.As(new(propertyv2.PropertyServiceServer))),
		fx.Annotate(handler.NewHostGroupV1Controller, fx.As(new(host_groupv1.HostGroupServiceServer))),
		fx.Annotate(handler.NewHostGroupV2Controller, fx.As(new(host_groupv2.HostGroupServiceServer))),
		fx.Annotate(handler.NewHostV1Controller, fx.As(new(hostv1.HostServiceServer))),
		fx.Annotate(handler.NewHostV2Controller, fx.As(new(hostv2.HostServiceServer))),
		fx.Annotate(handler.NewStaticConfigV1Controller, fx.As(new(configv1.StaticConfigServiceServer))),
		fx.Annotate(handler.NewStaticConfigV2Controller, fx.As(new(configv2.StaticConfigServiceServer))),
		fx.Annotate(handler.NewConfigV1Controller, fx.As(new(configv1.DynamicConfigServiceServer))),
		fx.Annotate(handler.NewConfigV2Controller, fx.As(new(configv2.DynamicConfigServiceServer))),
		fx.Annotate(handler.NewCompetitionV1Controller, fx.As(new(competitionv1.CompetitionServiceServer))),
		fx.Annotate(handler.NewCompetitionV2Controller, fx.As(new(competitionv2.CompetitionServiceServer))),
		fx.Annotate(handler.NewCheckV1Controller, fx.As(new(checkv1.CheckServiceServer))),
		fx.Annotate(handler.NewCheckV2Controller, fx.As(new(checkv2.CheckServiceServer))),
	),

	// Add them to the grpc server object
	fx.Invoke(
		checkv1.RegisterCheckServiceServer,
		checkv2.RegisterCheckServiceServer,
		competitionv1.RegisterCompetitionServiceServer,
		competitionv2.RegisterCompetitionServiceServer,
		configv1.RegisterStaticConfigServiceServer,
		configv2.RegisterStaticConfigServiceServer,
		configv1.RegisterDynamicConfigServiceServer,
		configv2.RegisterDynamicConfigServiceServer,
		hostv1.RegisterHostServiceServer,
		hostv2.RegisterHostServiceServer,
		host_groupv1.RegisterHostGroupServiceServer,
		host_groupv2.RegisterHostGroupServiceServer,
		propertyv1.RegisterPropertyServiceServer,
		propertyv2.RegisterPropertyServiceServer,
		reportv1.RegisterReportServiceServer,
		reportv2.RegisterReportServiceServer,
		roundv1.RegisterRoundServiceServer,
		roundv2.RegisterRoundServiceServer,
		servicev1.RegisterServiceServiceServer,
		servicev2.RegisterServiceServiceServer,
		service_groupv1.RegisterServiceGroupServiceServer,
		service_groupv2.RegisterServiceGroupServiceServer,
		teamv1.RegisterTeamServiceServer,
		teamv2.RegisterTeamServiceServer,
		userv1.RegisterUserServiceServer,
		userv2.RegisterUserServiceServer,
		authv1.RegisterAuthServiceServer,
		authv2.RegisterAuthServiceServer,
		policyv1.RegisterPolicyServiceServer,
		policyv2.RegisterPolicyServiceServer,
		healthv1.RegisterHealthServer,
	),
)
