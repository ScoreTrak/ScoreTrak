package cmd

import (
	"context"
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	authpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/auth/v1"
	v1 "github.com/ScoreTrak/ScoreTrak/pkg/proto/proto/v1"
	userpb "github.com/ScoreTrak/ScoreTrak/pkg/proto/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"log"

	"github.com/golang-jwt/jwt/v4"

	"github.com/spf13/cobra"
)

var adminName string
var adminPassword string
var newUserName string
var newUserPassword string
var address string
var teamId string

// createSuperUserCmd represents the createSuperUser command
var createSuperUserCmd = &cobra.Command{
	Use:   "createSuperUser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createSuperUser called")
		createNewSuperUser(address, adminName, adminPassword, newUserName, newUserPassword, teamId)
	},
}

func init() {
	rootCmd.AddCommand(createSuperUserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createSuperUserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createSuperUserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createSuperUserCmd.Flags().StringVarP(&adminName, "admin-name", "", "", "username of user with privileges")
	createSuperUserCmd.Flags().StringVarP(&adminPassword, "admin-password", "", "", "password of user with privileges")
	createSuperUserCmd.Flags().StringVarP(&newUserName, "new-user-name", "", "", "Username of the new user")
	createSuperUserCmd.Flags().StringVarP(&newUserPassword, "new-user-password", "", "", "Password of the new user")
	createSuperUserCmd.Flags().StringVarP(&address, "address", "", "", "Address of the scoretrak node")
	createSuperUserCmd.Flags().StringVarP(&teamId, "team-id", "", "", "ID of the team to which the user will be assigned, by default uses same team-id as admin user")
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
