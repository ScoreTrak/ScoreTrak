package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/ScoreTrak/ScoreTrak/pkg/auth"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/cobra"
	authv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/auth/v1"
	protov1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/proto/v1"
	userv1 "go.buf.build/grpc/go/scoretrak/scoretrakapis/scoretrak/user/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var adminName string
var adminPassword string
var newUserName string
var newUserPassword string
var address string
var teamID string

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
		log.Println("createSuperUser called")
		err := createNewSuperUser(address, adminName, adminPassword, newUserName, newUserPassword, teamID)
		if err != nil {
			log.Printf("%v", err)
		}
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
	createSuperUserCmd.Flags().StringVarP(&teamID, "team-id", "", "", "ID of the team to which the user will be assigned, by default uses same team-id as admin user")
}

func createNewSuperUser(address string, adminUsername string, adminPassword string, newUsername string, newPassword string, teamID string) error {
	cc, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient := authv1.NewAuthServiceClient(cc)
	resp, err := authClient.Login(context.Background(), &authv1.LoginRequest{Password: adminPassword, Username: adminUsername})
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
	userCli := userv1.NewUserServiceClient(cc)
	userResponse, err := userCli.GetByUsername(ctx, &userv1.GetByUsernameRequest{Username: newUsername})
	if err != nil || userResponse.GetUser() == nil || userResponse.GetUser().Username != newUsername {
		if teamID == "" {
			teamID = claims.TeamID
		}
		if _, err = userCli.Store(ctx, &userv1.StoreRequest{Users: []*userv1.User{{
			Username: newUsername,
			TeamId:   &protov1.UUID{Value: teamID},
			Password: newPassword,
			Role:     userv1.Role_ROLE_BLACK,
		}}}); err != nil {
			return fmt.Errorf("failed to create user: %w", err)
		}
	} else {
		log.Printf("User %s already exists", newUsername)
	}
	return nil
}
