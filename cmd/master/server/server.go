package server

import (
	"context"
	"errors"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	checkService "github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	competitionService "github.com/ScoreTrak/ScoreTrak/pkg/competition/competitionservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	configService "github.com/ScoreTrak/ScoreTrak/pkg/config/configservice"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	checkpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/check/v1"
	competitionpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/competition/v1"

	"github.com/ScoreTrak/ScoreTrak/pkg/di/util"
	hostService "github.com/ScoreTrak/ScoreTrak/pkg/host/hostservice"

	"github.com/ScoreTrak/ScoreTrak/pkg/handler"
	configpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/config/v1"
	hostpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/host/v1"
	policypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/policy/v1"
	teampb "github.com/ScoreTrak/ScoreTrak/pkg/proto/team/v1"

	hostGroupService "github.com/ScoreTrak/ScoreTrak/pkg/hostgroup/hostgroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy/policyclient"
	policyService "github.com/ScoreTrak/ScoreTrak/pkg/policy/policyservice"
	host_grouppb "github.com/ScoreTrak/ScoreTrak/pkg/proto/host_group/v1"

	propertyService "github.com/ScoreTrak/ScoreTrak/pkg/property/propertyservice"
	propertypb "github.com/ScoreTrak/ScoreTrak/pkg/proto/property/v1"

	reportpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/report/v1"
	roundpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/round/v1"
	servicepb "github.com/ScoreTrak/ScoreTrak/pkg/proto/service/v1"
	service_grouppb "github.com/ScoreTrak/ScoreTrak/pkg/proto/service_group/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/queue"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/report/reportclient"
	reportService "github.com/ScoreTrak/ScoreTrak/pkg/report/reportservice"
	roundService "github.com/ScoreTrak/ScoreTrak/pkg/round/roundservice"
	serviceService "github.com/ScoreTrak/ScoreTrak/pkg/service/serviceservice"
	serviceGroupService "github.com/ScoreTrak/ScoreTrak/pkg/servicegroup/servicegroupservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	teamService "github.com/ScoreTrak/ScoreTrak/pkg/team/teamservice"

	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	userpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/user/v1"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	userService "github.com/ScoreTrak/ScoreTrak/pkg/user/userservice"
	"github.com/gofrs/uuid"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/jackc/pgconn"
	"go.uber.org/dig"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

var ErrProdCertMissing = errors.New("production requires certfile, and keyfile")

func Start(staticConfig config.StaticConfig, d *dig.Container, db *gorm.DB) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", staticConfig.Port))
	if err != nil {
		log.Panicf("Failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	if staticConfig.CertFile != "" && staticConfig.KeyFile != "" {
		creds, sslErr := credentials.NewClientTLSFromFile(staticConfig.CertFile, staticConfig.KeyFile)
		if sslErr != nil {
			return err
		}
		opts = append(opts, grpc.Creds(creds))
	} else if config.GetStaticConfig().Prod {
		return ErrProdCertMissing
	}

	var middlewareChainsUnary []grpc.UnaryServerInterceptor
	var middlewareChainsStream []grpc.StreamServerInterceptor

	// Logging & Recovery
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

	// Recovery
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

	// Load policy into DB & Create Policy service
	var policyServ policyService.Serv
	if err := d.Invoke(func(s policyService.Serv) {
		policyServ = s
	}); err != nil {
		return err
	}
	p := &policy.Policy{ID: 1}
	err = db.Create(p).Error
	if err != nil {
		var serr *pgconn.PgError
		ok := errors.As(err, &serr)
		if !ok {
			if serr.Code != "23505" {
				panic(err)
			} else {
				db.Take(p)
			}
		}
	}

	// Repo Store
	repoStore := util.NewStore()

	// Authorization And Authentication Middleware
	jwtManager := auth.NewJWTManager(config.GetJWTConfig().Secret, time.Duration(config.GetJWTConfig().TimeoutInSeconds)*time.Second)
	var pubsub queue.MasterStreamPubSub

	pubsub, err = queue.NewMasterStreamPubSub(staticConfig.Queue)
	if err != nil {
		return err
	}

	policyClient := policyclient.NewPolicyClient(p, staticConfig.PubSubConfig, repoStore.Policy, pubsub)
	go func() {
		policyClient.PolicyClient()
	}()

	{
		ai := auth.NewAuthInterceptor(jwtManager, policyClient)
		middlewareChainsUnary = append(middlewareChainsUnary, ai.Unary())
		middlewareChainsStream = append(middlewareChainsStream, ai.Stream())
	}

	// Middleware Chaining
	{
		opts = append(opts, grpc_middleware.WithUnaryServerChain(middlewareChainsUnary...))
		opts = append(opts, grpc_middleware.WithStreamServerChain(middlewareChainsStream...))
	}

	// Middleware Chaining
	// {
	// 	opts = append(opts, grpc.MaxRecvMsgSize()) //Todo: Figure out a way to pass parameter for MaxRecvMsgSize (Can help with large competition files)
	// }
	// {
	// 	opts = append(opts)
	// }

	s := grpc.NewServer(opts...)

	// Routes
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

			var staticConfigSvc configService.StaticServ
			if err := d.Invoke(func(s configService.StaticServ) {
				staticConfigSvc = s
			}); err != nil {
				return err
			}
			configpb.RegisterStaticConfigServiceServer(s, handler.NewStaticConfigController(staticConfigSvc))
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
					var serr *pgconn.PgError
					ok := errors.As(err, &serr)
					if !ok || serr.Code != "23505" {
						return err
					}
				}
			}
			reportClient := reportclient.NewReportClient(staticConfig.PubSubConfig, repoStore.Report, pubsub)
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
			fls := false
			idx := uint64(0)
			err := teamSvc.Store(context.Background(), []*team.Team{
				{
					ID:    uuid1,
					Name:  "Black Team",
					Pause: &fls,
					Index: &idx,
					Users: nil,
					Hosts: nil,
					Hide:  &fls,
				},
			})
			if err != nil {
				var serr *pgconn.PgError
				ok := errors.As(err, &serr)
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
			passwordHash, _ := bcrypt.GenerateFromPassword([]byte("changeme"), bcrypt.DefaultCost)
			err := userServ.Store(context.Background(), []*user.User{{
				ID:           uuid1,
				Role:         user.Black,
				TeamID:       uuid1,
				Username:     "admin",
				PasswordHash: string(passwordHash),
			}})
			if err != nil {
				var serr *pgconn.PgError
				ok := errors.As(err, &serr)
				if !ok || serr.Code != "23505" {
					return err
				}
			}
			userpb.RegisterUserServiceServer(s, handler.NewUserController(userServ, policyClient))

			authpb.RegisterAuthServiceServer(s, handler.NewAuthController(userServ, jwtManager))
		}

		{
			policypb.RegisterPolicyServiceServer(s, handler.NewPolicyController(policyServ, policyClient))
		}
	}

	if !staticConfig.Prod {
		reflection.Register(s)
	}

	go func() {
		log.Println("Starting Server...")
		if err := s.Serve(lis); err != nil {
			time.Sleep(time.Second)
			log.Panicf("failed to serve: %v", err)
		}
	}()

	waitForSignal()
	if err := lis.Close(); err != nil {
		log.Panicf("Error on closing the listener : %v", err)
	}
	s.Stop()
	log.Println("Bye!")
	return nil
}

func waitForSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs
}
