package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	v1 "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	userpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/user/v1"
	"github.com/golang-jwt/jwt/v4"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "create-super-user",
				Usage: "create new super user",
				Action: func(c *cli.Context) error {
					return createNewSuperUser(c.String("address"), c.String("admin-name"), c.String("admin-password"), c.String("new-user-name"), c.String("new-user-password"), c.String("team-id"))
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "admin-name",
						Value: "admin",
						Usage: "Username of user with privileges",
					},
					&cli.StringFlag{
						Name:     "admin-password",
						Usage:    "Username of user with privileges",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "new-user-name",
						Usage:    "Username of the new user",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "new-user-password",
						Usage:    "Password of the new User",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "address",
						Usage:    "Address of the ScoreTrak node",
						Required: true,
					},
					&cli.StringFlag{
						Name:  "team-id",
						Value: "",
						Usage: "ID of the team to which the user will be assigned, by default uses same team-id as admin user",
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func createNewSuperUser(address string, adminUsername string, adminPassword string, newUsername string, newPassword string, teamID string) error {
	cc, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient := authpb.NewAuthServiceClient(cc)
	resp, err := authClient.Login(context.Background(), &authpb.LoginRequest{Password: adminPassword, Username: adminUsername})
	if err != nil {
		return err
	}
	token, _, err := new(jwt.Parser).ParseUnverified(resp.AccessToken, &auth.UserClaims{})
	if err != nil {
		panic(err)
	}
	claims, ok := token.Claims.(*auth.UserClaims)
	if !ok {
		log.Fatalf("invalid token claims")
	}
	log.Println("Authentication token is valid until " + claims.ExpiresAt.Time.String())

	ctx := metadata.NewOutgoingContext(context.Background(), metadata.Pairs("authorization", resp.AccessToken))
	userCli := userpb.NewUserServiceClient(cc)
	userResponse, err := userCli.GetByUsername(ctx, &userpb.GetByUsernameRequest{Username: newUsername})
	if err != nil || userResponse.GetUser() == nil || userResponse.GetUser().Username != newUsername {
		if teamID == "" {
			teamID = claims.TeamID
		}
		if _, err = userCli.Store(ctx, &userpb.StoreRequest{Users: []*userpb.User{{
			Username: newUsername,
			TeamId:   &v1.UUID{Value: teamID},
			Password: newPassword,
			Role:     userpb.Role_ROLE_BLACK,
		}}}); err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		log.Printf("User %s already exists", newUsername)
	}
	return nil
}
