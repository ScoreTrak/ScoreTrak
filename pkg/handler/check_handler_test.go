package handler

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/ScoreTrak/ScoreTrak/pkg/check/checkservice"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	. "github.com/ScoreTrak/ScoreTrak/pkg/config/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/orm"
	. "github.com/ScoreTrak/ScoreTrak/pkg/storage/orm/testutil"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/golang-jwt/jwt/v4"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	. "github.com/smartystreets/goconvey/convey"
	checkv1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/check/v1"
	protov1 "go.buf.build/library/go-grpc/scoretrak/scoretrakapis/scoretrak/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"log"
	"net"
	"os"
	"testing"
)

func TestCheckSpec(t *testing.T) {
	var c config.StaticConfig
	autoTest := os.Getenv("AUTO_TEST")
	if autoTest == "TRUE" {
		c = NewTestConfigClone("../../configs/test-config.yml")
	} else {
		c = NewConfigClone(SetupConfig("dev-config.yml"))
	}
	c.DB.Cockroach.Database = "scoretrak_test_handler_check"
	db := SetupCockroachDB(c.DB)
	t.Parallel() //t.Parallel should be placed after SetupCockroachDB because gorm has race conditions on Hook register
	err := util.CreateAllTables(db)
	if err != nil {
		panic(err)
	}

	userClaims := []*auth.UserClaims{
		{
			RegisteredClaims: jwt.RegisteredClaims{},
			Username:         "TeamOneUser",
			TeamID:           "11111111-1111-1111-1111-111111111111",
			Role:             "black",
		},
		{
			RegisteredClaims: jwt.RegisteredClaims{},
			Username:         "TeamTwoUser1",
			TeamID:           "22222222-2222-2222-2222-222222222222",
			Role:             "blue",
		},
		{
			RegisteredClaims: jwt.RegisteredClaims{},
			Username:         "TeamTwoUser2",
			TeamID:           "22222222-2222-2222-2222-222222222222",
			Role:             "red",
		},
	}

	for _, claim := range userClaims {
		Convey("Create Tables and Load Data", t, func() {
			DataPreload(db)
			Convey("Creating Round, Service and Check repos, services, ", func() {
				cr := orm.NewCheckRepo(db)
				cs := checkservice.NewCheckServ(cr)
				us := util.Store{
					Round:        orm.NewRoundRepo(db),
					Service:      orm.NewServiceRepo(db),
					Check:        cr,
					Host:         orm.NewHostRepo(db),
					HostGroup:    orm.NewHostGroupRepo(db),
					ServiceGroup: orm.NewServiceGroupRepo(db),
					Team:         orm.NewTeamRepo(db),
					Property:     orm.NewPropertyRepo(db),
					Config:       orm.NewConfigRepo(db),
					Report:       orm.NewReportRepo(db),
					Policy:       orm.NewPolicyRepo(db),
					Users:        orm.NewUserRepo(db),
				}
				const bufSize = 1024 * 1024
				lis := bufconn.Listen(bufSize)
				var opts []grpc.ServerOption

				var middlewareChainsUnary []grpc.UnaryServerInterceptor
				var middlewareChainsStream []grpc.StreamServerInterceptor

				middlewareChainsUnary = append(middlewareChainsUnary, func(
					ctx context.Context,
					req interface{},
					info *grpc.UnaryServerInfo,
					handler grpc.UnaryHandler,
				) (interface{}, error) {
					ctx = context.WithValue(ctx, auth.KeyClaim, claim)
					return handler(ctx, req)
				})

				middlewareChainsStream = append(middlewareChainsStream, func(
					srv interface{},
					stream grpc.ServerStream,
					info *grpc.StreamServerInfo,
					handler grpc.StreamHandler,
				) error {
					return handler(srv, auth.StreamClaimInjector{ServerStream: stream, Claims: claim})
				})

				opts = append(opts, grpc_middleware.WithUnaryServerChain(middlewareChainsUnary...))
				opts = append(opts, grpc_middleware.WithStreamServerChain(middlewareChainsStream...))

				s := grpc.NewServer(opts...)
				checkv1.RegisterCheckServiceServer(s, NewCheckController(cs, &us))

				go func() {
					if err := s.Serve(lis); err != nil {
						log.Panicf("Server exited with error: %v", err)
					}
				}()

				ctx := context.Background()
				conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
					return lis.Dial()
				}), grpc.WithInsecure())
				if err != nil {
					t.Fatalf("Failed to dial bufnet: %v", err)
				}
				defer conn.Close()
				client := checkv1.NewCheckServiceClient(conn)

				Convey("Then run GetAllByRoundID", func() {

				})

				Convey("Then run GetAllByServiceID", func() {

				})

				Convey("Then run GetByRoundServiceID", func() {
					Convey("With Correct parameters", func() {
						if claim.Role == user.Black {
							resp, err := client.GetByRoundServiceID(ctx, &checkv1.GetByRoundServiceIDRequest{
								ServiceId: &protov1.UUID{Value: "11111111-1111-1111-1111-111111111111"},
								RoundId:   1,
							})
							So(err, ShouldBeNil)
							So(resp.Check, ShouldNotBeNil)
							So(resp.Check.Log, ShouldEqual, "Successful Check One!")
						}
						if claim.Role == user.Blue {
							resp, err := client.GetByRoundServiceID(ctx, &checkv1.GetByRoundServiceIDRequest{
								ServiceId: &protov1.UUID{Value: "22222222-2222-2222-2222-222222222222"},
								RoundId:   2,
							})
							So(err, ShouldBeNil)
							So(resp.Check, ShouldNotBeNil)
							So(resp.Check.Log, ShouldEqual, "Successful Check Two!")
							resp, err = client.GetByRoundServiceID(ctx, &checkv1.GetByRoundServiceIDRequest{
								ServiceId: &protov1.UUID{Value: "11111111-1111-1111-1111-111111111111"},
								RoundId:   1,
							})
							So(err, ShouldNotBeNil)
							So(resp, ShouldBeNil)
							e, ok := status.FromError(err)
							So(ok, ShouldBeTrue)
							So(e.Code(), ShouldEqual, codes.PermissionDenied)
							So(resp, ShouldBeNil)
						}

					})

					Convey("With following incorrect parameters", func() {
						Convey("ServiceId", func() {
							Convey("Set to non existing ID", func() {
								resp, err := client.GetByRoundServiceID(ctx, &checkv1.GetByRoundServiceIDRequest{
									ServiceId: &protov1.UUID{Value: "55555555-5555-5555-5555-555555555500"},
									RoundId:   1,
								})
								So(err, ShouldNotBeNil)
								e, ok := status.FromError(err)
								So(ok, ShouldBeTrue)
								fmt.Println(e.Err())
								So(e.Code(), ShouldEqual, codes.NotFound)
								So(resp, ShouldBeNil)
							})
							Convey("Set to null", func() {
								resp, err := client.GetByRoundServiceID(ctx, &checkv1.GetByRoundServiceIDRequest{
									ServiceId: nil,
									RoundId:   1,
								})
								So(err, ShouldNotBeNil)
								e, ok := status.FromError(err)
								So(ok, ShouldBeTrue)
								fmt.Println(e.Err())
								So(e.Code(), ShouldEqual, codes.InvalidArgument)
								So(resp, ShouldBeNil)
							})

							Convey("Set null uuid", func() {
								resp, err := client.GetByRoundServiceID(ctx, &checkv1.GetByRoundServiceIDRequest{
									ServiceId: &protov1.UUID{},
									RoundId:   1,
								})
								So(err, ShouldNotBeNil)
								e, ok := status.FromError(err)
								So(ok, ShouldBeTrue)
								fmt.Println(e.Err())
								So(e.Code(), ShouldEqual, codes.InvalidArgument)
								So(resp, ShouldBeNil)
							})

						})
						Convey("RoundId", func() {

						})
					})
				})
			})

			Reset(func() {
				TruncateAllTables(db)
			})
		})
	}

	DropDB(db, c)

}
