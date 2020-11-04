package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	checkService "github.com/ScoreTrak/ScoreTrak/pkg/check/check_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkpb"
	competitionService "github.com/ScoreTrak/ScoreTrak/pkg/competition/competition_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	configService "github.com/ScoreTrak/ScoreTrak/pkg/config/config_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/config/configpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/di/util"
	hostService "github.com/ScoreTrak/ScoreTrak/pkg/host/host_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host/hostpb"
	hostGroupService "github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_group_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/host_group/host_grouppb"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_client"
	policyService "github.com/ScoreTrak/ScoreTrak/pkg/policy/policy_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policypb"
	propertyService "github.com/ScoreTrak/ScoreTrak/pkg/property/property_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/property/propertypb"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/handler"
	"github.com/ScoreTrak/ScoreTrak/pkg/proto/utilpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/report_client"
	reportService "github.com/ScoreTrak/ScoreTrak/pkg/report/report_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportpb"
	"github.com/ScoreTrak/ScoreTrak/pkg/role"
	roundService "github.com/ScoreTrak/ScoreTrak/pkg/round/round_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/round/roundpb"
	serviceService "github.com/ScoreTrak/ScoreTrak/pkg/service/service_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service/servicepb"
	serviceGroupService "github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_group_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/service_group/service_grouppb"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	teamService "github.com/ScoreTrak/ScoreTrak/pkg/team/team_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/team/teampb"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	userService "github.com/ScoreTrak/ScoreTrak/pkg/user/user_service"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userpb"
	"github.com/gofrs/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/jackc/pgconn"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(staticConfig config.StaticConfig, d *dig.Container, db *gorm.DB) error {
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
			middlewareChainsUnary = append(middlewareChainsUnary, grpc_recovery.UnaryServerInterceptor(recoveryOpts...))
			middlewareChainsStream = append(middlewareChainsStream, grpc_recovery.StreamServerInterceptor(recoveryOpts...))
		}
	}

	//Load policy into DB & Create Policy service
	var policyServ policyService.Serv
	if err := d.Invoke(func(s policyService.Serv) {
		policyServ = s
	}); err != nil {
		return err
	}
	p := &policy.Policy{ID: 1}
	err = db.Create(p).Error
	if err != nil {
		serr, ok := err.(*pgconn.PgError)
		if !ok {
			if serr.Code != "23505" {
				panic(err)
			} else {
				db.Take(p)
			}
		}
	}

	//Repo Store
	repoStore := util.NewStore()

	//Authorization And Authentication Middleware
	jwtManager := auth.NewJWTManager(config.GetJWTConfig().Secret, time.Duration(config.GetJWTConfig().TimeoutInSeconds)*time.Second)
	var pubsub queue.MasterStreamPubSub
	if config.GetQueueConfig().Use != "none" {
		pubsub, err = queue.NewMasterStreamPubSub(staticConfig.Queue)
		if err != nil {
			return err
		}
	}

	policyClient := policy_client.NewPolicyClient(p, staticConfig.PubSubConfig, repoStore.Policy, pubsub)
	go func() {
		policyClient.PolicyClient()
	}()

	{
		ai := auth.NewAuthInterceptor(jwtManager, policyClient)
		middlewareChainsUnary = append(middlewareChainsUnary, ai.Unary())
		middlewareChainsStream = append(middlewareChainsStream, ai.Stream())
	}

	//Middleware Chaining
	{
		opts = append(opts, grpc_middleware.WithUnaryServerChain(middlewareChainsUnary...))
		opts = append(opts, grpc_middleware.WithStreamServerChain(middlewareChainsStream...))
	}

	{
		opts = append(opts)
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
			checkpb.RegisterCheckServiceServer(s, handler.NewCheckController(checkSvc, repoStore))
		}
		{
			comptSvc := competitionService.NewCompetitionServ(repoStore)
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
			hostpb.RegisterHostServiceServer(s, handler.NewHostController(hostSvc, repoStore))
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
			propertypb.RegisterPropertyServiceServer(s, handler.NewPropertyController(propertySvc, repoStore))
		}
		{
			var reportSvc reportService.Serv
			if err := d.Invoke(func(s reportService.Serv) {
				reportSvc = s
			}); err != nil {
				return err
			}

			var count int64
			if count != 1 {
				err := db.Create(report.NewReport()).Error
				if err != nil {
					serr, ok := err.(*pgconn.PgError)
					if !ok || serr.Code != "23505" {
						return err
					}
				}
			}
			reportClient := report_client.NewReportClient(staticConfig.PubSubConfig, repoStore.Report, pubsub)
			go func() {
				reportClient.ReportClient()
			}()
			reportpb.RegisterReportServiceServer(s, handler.NewReportController(reportSvc, reportClient, policyClient))
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
			servicepb.RegisterServiceServiceServer(s, handler.NewServiceController(serviceSvc, repoStore))
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

		var uuid1 = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")
		{
			var teamSvc teamService.Serv
			if err := d.Invoke(func(s teamService.Serv) {
				teamSvc = s
			}); err != nil {
				return err
			}
			tru := true
			idx := uint64(0)
			err := teamSvc.Store(context.Background(), []*team.Team{
				{
					ID:      uuid1,
					Name:    "Black Team",
					Enabled: &tru,
					Index:   &idx,
					Users:   nil,
					Hosts:   nil,
				},
			})
			if err != nil {
				serr, ok := err.(*pgconn.PgError)
				if !ok || serr.Code != "23505" {
					return err
				}
			}
			teampb.RegisterTeamServiceServer(s, handler.NewTeamController(teamSvc))
		}

		{
			var userServ userService.Serv
			if err := d.Invoke(func(s userService.Serv) {
				userServ = s
			}); err != nil {
				return err
			}
			usr, _ := handler.ConvertUserPBtoUser(true, &userpb.User{
				Id:       &utilpb.UUID{Value: uuid1.String()},
				Username: "admin",
				TeamId:   &utilpb.UUID{Value: uuid1.String()},
				Password: "changeme",
				Role:     handler.UserRoleToRolePB(role.Black),
			})
			err := userServ.Store(context.Background(), []*user.User{usr})
			if err != nil {
				serr, ok := err.(*pgconn.PgError)
				if !ok || serr.Code != "23505" {
					return err
				}
			}
			userpb.RegisterUserServiceServer(s, handler.NewUserController(userServ))

			auth.RegisterAuthServiceServer(s, handler.NewAuthController(userServ, jwtManager))
		}

		{
			policypb.RegisterPolicyServiceServer(s, handler.NewPolicyController(policyServ, policyClient))
		}

	}

	if !staticConfig.Prod {
		reflection.Register(s)
	}

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			time.Sleep(time.Second)
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
