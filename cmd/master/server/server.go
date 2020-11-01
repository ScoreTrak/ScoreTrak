package server

import (
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	checkService "github.com/ScoreTrak/ScoreTrak/pkg/check/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionpb"
	competitionService "github.com/ScoreTrak/ScoreTrak/pkg/competition/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configpb"
	configService "github.com/ScoreTrak/ScoreTrak/pkg/config/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostpb"
	hostService "github.com/ScoreTrak/ScoreTrak/pkg/host/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_grouppb"
	hostGroupService "github.com/ScoreTrak/ScoreTrak/pkg/host_group/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertypb"
	propertyService "github.com/ScoreTrak/ScoreTrak/pkg/property/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/handler"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	reportService "github.com/ScoreTrak/ScoreTrak/pkg/report/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundpb"
	roundService "github.com/ScoreTrak/ScoreTrak/pkg/round/service"
	serviceService "github.com/ScoreTrak/ScoreTrak/pkg/service/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicepb"
	serviceGroupService "github.com/ScoreTrak/ScoreTrak/pkg/service_group/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_grouppb"
	teamService "github.com/ScoreTrak/ScoreTrak/pkg/team/service"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teampb"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Start(staticConfig config.StaticConfig, d *dig.Container) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", staticConfig.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if staticConfig.CertFile != "" && staticConfig.KeyFile != "" {
		creds, sslErr := credentials.NewClientTLSFromFile(staticConfig.CertFile, staticConfig.KeyFile)
		if sslErr != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	} else if config.GetStaticConfig().Prod {
		return errors.New("you must specify certfile, and keyfile when running in production")
	}

	var middlewareChainsUnary []grpc.UnaryServerInterceptor
	var middlewareChainsStream []grpc.StreamServerInterceptor

	//Logging & Recovery
	{
		var zapLogger *zap.Logger
		if staticConfig.Prod {
			zapLogger, err = zap.NewProduction()
			if err != nil {
				return err
			}

		} else {
			zapLogger, err = zap.NewDevelopment()
			if err != nil {
				return err
			}
		}
		customFunc := grpc_zap.DefaultCodeToLevel
		grpc_zap.ReplaceGrpcLoggerV2(zapLogger)
		logOpts := []grpc_zap.Option{
			grpc_zap.WithLevels(customFunc),
		}

		middlewareChainsUnary = append(middlewareChainsUnary, []grpc.UnaryServerInterceptor{
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.UnaryServerInterceptor(zapLogger, logOpts...),
		}...)

		middlewareChainsStream = append(middlewareChainsStream, []grpc.StreamServerInterceptor{
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_zap.StreamServerInterceptor(zapLogger, logOpts...),
		}...)

	}

	//Recovery
	{
		customFunc := func(p interface{}) (err error) {
			return status.Errorf(codes.Unknown, "panic triggered: %v", p)
		}

		recoveryOpts := []grpc_recovery.Option{
			grpc_recovery.WithRecoveryHandler(customFunc),
		}
		if staticConfig.Prod {

			middlewareChainsUnary = append(middlewareChainsUnary, []grpc.UnaryServerInterceptor{
				grpc_recovery.UnaryServerInterceptor(recoveryOpts...),
			}...)

			middlewareChainsStream = append(middlewareChainsStream, []grpc.StreamServerInterceptor{
				grpc_recovery.StreamServerInterceptor(recoveryOpts...),
			}...)
		}
	}
	//Middleware Chaining
	{
		opts = append(opts, grpc_middleware.WithUnaryServerChain(middlewareChainsUnary...))
		opts = append(opts, grpc_middleware.WithStreamServerChain(middlewareChainsStream...))
	}

	s := grpc.NewServer(opts...)

	//Routes
	{
		{
			var checkSvc checkService.Serv
			if err := d.Invoke(func(s checkService.Serv) {
				checkSvc = s
			}); err != nil {
				return err
			}
			checkpb.RegisterCheckServiceServer(s, handler.NewCheckController(checkSvc))
		}
		{
			comptSvc := competitionService.NewCompetitionServ(util.NewStore())
			competitionpb.RegisterCompetitionServiceServer(s, handler.NewCompetitionController(comptSvc))
		}
		{
			var configSvc configService.Serv
			if err := d.Invoke(func(s configService.Serv) {
				configSvc = s
			}); err != nil {
				return err
			}
			configpb.RegisterDynamicConfigServiceServer(s, handler.NewConfigController(configSvc))
		}
		{
			var hostSvc hostService.Serv
			if err := d.Invoke(func(s hostService.Serv) {
				hostSvc = s
			}); err != nil {
				return err
			}
			hostpb.RegisterHostServiceServer(s, handler.NewHostController(hostSvc))
		}
		{
			var hostGroupSvc hostGroupService.Serv
			if err := d.Invoke(func(s hostGroupService.Serv) {
				hostGroupSvc = s
			}); err != nil {
				return err
			}
			host_grouppb.RegisterHostGroupServiceServer(s, handler.NewHostGroupController(hostGroupSvc))
		}
		{
			var propertySvc propertyService.Serv
			if err := d.Invoke(func(s propertyService.Serv) {
				propertySvc = s
			}); err != nil {
				return err
			}
			propertypb.RegisterPropertyServiceServer(s, handler.NewPropertyController(propertySvc))
		}
		{
			var reportSvc reportService.Serv
			if err := d.Invoke(func(s reportService.Serv) {
				reportSvc = s
			}); err != nil {
				return err
			}
			reportpb.RegisterReportServiceServer(s, handler.NewReportController(reportSvc))
		}
		{
			var roundSvc roundService.Serv
			if err := d.Invoke(func(s roundService.Serv) {
				roundSvc = s
			}); err != nil {
				return err
			}
			roundpb.RegisterRoundServiceServer(s, handler.NewRoundController(roundSvc))
		}
		{
			var serviceSvc serviceService.Serv
			if err := d.Invoke(func(s serviceService.Serv) {
				serviceSvc = s
			}); err != nil {
				return err
			}
			servicepb.RegisterServiceServiceServer(s, handler.NewServiceController(serviceSvc))
		}

		{
			var serviceGroupSvc serviceGroupService.Serv
			if err := d.Invoke(func(s serviceGroupService.Serv) {
				serviceGroupSvc = s
			}); err != nil {
				return err
			}
			service_grouppb.RegisterServiceGroupServiceServer(s, handler.NewServiceGroupController(serviceGroupSvc))
		}

		{
			var teamSvc teamService.Serv
			if err := d.Invoke(func(s teamService.Serv) {
				teamSvc = s
			}); err != nil {
				return err
			}
			teampb.RegisterTeamServiceServer(s, handler.NewTeamController(teamSvc))
		}
	}

	if !staticConfig.Prod {
		reflection.Register(s)
	}

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	waitForSignal()
	if err := lis.Close(); err != nil {
		log.Fatalf("Error on closing the listener : %v", err)
	}
	s.Stop()
	fmt.Println("Bye!")
	return nil
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}
