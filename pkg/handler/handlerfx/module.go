package handlerfx

import (
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/grpc/health/v1/healthv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/auth/v1/authv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/auth/v2/authv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/check/v1/checkv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/check/v2/checkv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/competition/v1/competitionv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/competition/v2/competitionv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/config/v1/configv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/config/v2/configv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/host/v1/hostv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/host/v2/hostv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/host_group/v1/host_groupv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/host_group/v2/host_groupv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/policy/v1/policyv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/policy/v2/policyv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/property/v1/propertyv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/property/v2/propertyv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/report/v1/reportv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/report/v2/reportv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/round/v1/roundv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/round/v2/roundv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/service/v1/servicev1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/service/v2/servicev2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/service_group/v1/service_groupv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/service_group/v2/service_groupv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/team/v1/teamv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/team/v2/teamv2grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/user/v1/userv1grpc"
	"buf.build/gen/go/scoretrak/scoretrakapis/grpc/go/scoretrak/user/v2/userv2grpc"
	"github.com/ScoreTrak/ScoreTrak/pkg/handler"
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
		fx.Annotate(handler.NewHealthV1Controller, fx.As(new(healthv1grpc.HealthServer))),
		fx.Annotate(handler.NewPolicyV1Controller, fx.As(new(policyv1grpc.PolicyServiceServer))),
		fx.Annotate(handler.NewPolicyV2Controller, fx.As(new(policyv2grpc.PolicyServiceServer))),
		fx.Annotate(handler.NewAuthV1Controller, fx.As(new(authv1grpc.AuthServiceServer))),
		fx.Annotate(handler.NewAuthV2Controller, fx.As(new(authv2grpc.AuthServiceServer))),
		fx.Annotate(handler.NewUserV1Controller, fx.As(new(userv1grpc.UserServiceServer))),
		fx.Annotate(handler.NewUserV2Controller, fx.As(new(userv2grpc.UserServiceServer))),
		fx.Annotate(handler.NewTeamV1Controller, fx.As(new(teamv1grpc.TeamServiceServer))),
		fx.Annotate(handler.NewTeamV2Controller, fx.As(new(teamv2grpc.TeamServiceServer))),
		fx.Annotate(handler.NewServiceV1Controller, fx.As(new(servicev1grpc.ServiceServiceServer))),
		fx.Annotate(handler.NewServiceV2Controller, fx.As(new(servicev2grpc.ServiceServiceServer))),
		fx.Annotate(handler.NewServiceGroupV1Controller, fx.As(new(service_groupv1grpc.ServiceGroupServiceServer))),
		fx.Annotate(handler.NewServiceGroupV2Controller, fx.As(new(service_groupv2grpc.ServiceGroupServiceServer))),
		fx.Annotate(handler.NewRoundV1Controller, fx.As(new(roundv1grpc.RoundServiceServer))),
		fx.Annotate(handler.NewRoundV2Controller, fx.As(new(roundv2grpc.RoundServiceServer))),
		fx.Annotate(handler.NewReportV1Controller, fx.As(new(reportv1grpc.ReportServiceServer))),
		fx.Annotate(handler.NewReportV2Controller, fx.As(new(reportv2grpc.ReportServiceServer))),
		fx.Annotate(handler.NewPropertyV1Controller, fx.As(new(propertyv1grpc.PropertyServiceServer))),
		fx.Annotate(handler.NewPropertyV2Controller, fx.As(new(propertyv2grpc.PropertyServiceServer))),
		fx.Annotate(handler.NewHostGroupV1Controller, fx.As(new(host_groupv1grpc.HostGroupServiceServer))),
		fx.Annotate(handler.NewHostGroupV2Controller, fx.As(new(host_groupv2grpc.HostGroupServiceServer))),
		fx.Annotate(handler.NewHostV1Controller, fx.As(new(hostv1grpc.HostServiceServer))),
		fx.Annotate(handler.NewHostV2Controller, fx.As(new(hostv2grpc.HostServiceServer))),
		fx.Annotate(handler.NewStaticConfigV1Controller, fx.As(new(configv1grpc.StaticConfigServiceServer))),
		fx.Annotate(handler.NewStaticConfigV2Controller, fx.As(new(configv2grpc.StaticConfigServiceServer))),
		fx.Annotate(handler.NewConfigV1Controller, fx.As(new(configv1grpc.DynamicConfigServiceServer))),
		fx.Annotate(handler.NewConfigV2Controller, fx.As(new(configv2grpc.DynamicConfigServiceServer))),
		fx.Annotate(handler.NewCompetitionV1Controller, fx.As(new(competitionv1grpc.CompetitionServiceServer))),
		fx.Annotate(handler.NewCompetitionV2Controller, fx.As(new(competitionv2grpc.CompetitionServiceServer))),
		fx.Annotate(handler.NewCheckV1Controller, fx.As(new(checkv1grpc.CheckServiceServer))),
		fx.Annotate(handler.NewCheckV2Controller, fx.As(new(checkv2grpc.CheckServiceServer))),
	),

	// Add them to the grpc server object
	fx.Invoke(
		checkv1grpc.RegisterCheckServiceServer,
		checkv2grpc.RegisterCheckServiceServer,
		competitionv1grpc.RegisterCompetitionServiceServer,
		competitionv2grpc.RegisterCompetitionServiceServer,
		configv1grpc.RegisterStaticConfigServiceServer,
		configv2grpc.RegisterStaticConfigServiceServer,
		configv1grpc.RegisterDynamicConfigServiceServer,
		configv2grpc.RegisterDynamicConfigServiceServer,
		hostv1grpc.RegisterHostServiceServer,
		hostv2grpc.RegisterHostServiceServer,
		host_groupv1grpc.RegisterHostGroupServiceServer,
		host_groupv2grpc.RegisterHostGroupServiceServer,
		propertyv1grpc.RegisterPropertyServiceServer,
		propertyv2grpc.RegisterPropertyServiceServer,
		reportv1grpc.RegisterReportServiceServer,
		reportv2grpc.RegisterReportServiceServer,
		roundv1grpc.RegisterRoundServiceServer,
		roundv2grpc.RegisterRoundServiceServer,
		servicev1grpc.RegisterServiceServiceServer,
		servicev2grpc.RegisterServiceServiceServer,
		service_groupv1grpc.RegisterServiceGroupServiceServer,
		service_groupv2grpc.RegisterServiceGroupServiceServer,
		teamv1grpc.RegisterTeamServiceServer,
		teamv2grpc.RegisterTeamServiceServer,
		userv1grpc.RegisterUserServiceServer,
		userv2grpc.RegisterUserServiceServer,
		authv1grpc.RegisterAuthServiceServer,
		authv2grpc.RegisterAuthServiceServer,
		policyv1grpc.RegisterPolicyServiceServer,
		policyv2grpc.RegisterPolicyServiceServer,
		healthv1grpc.RegisterHealthServer,
	),
)
