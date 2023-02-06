package seed

import (
	"context"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/policy"
	"github.com/ScoreTrak/ScoreTrak/pkg/report"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// create black team
// create admin user
// create config
// create policy
// create report

func DefaultSeed(staticConfig config.StaticConfig, dynamicConfig *config.DynamicConfig, repos *util.Store) error {
	ctx := context.Background()

	// CREATE BLACK TEAM
	var blackTeamUUID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")
	fls := false
	idx := uint64(0)
	err := repos.Team.Store(ctx, []*team.Team{{
		ID:    blackTeamUUID,
		Name:  "Black Team",
		Pause: &fls,
		Hide:  &fls,
		Index: &idx,
		Users: nil,
		Hosts: nil,
	}})
	if err != nil {
		return err
	}
	log.Println("Created black team")

	// CREATE ADMIN USER
	var adminUUID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(staticConfig.AdminPassword), bcrypt.DefaultCost)
	err = repos.Users.Store(context.Background(), []*user.User{{
		ID:           adminUUID,
		Role:         user.Black,
		TeamID:       blackTeamUUID,
		Username:     staticConfig.AdminUsername,
		PasswordHash: string(passwordHash),
	}})
	if err != nil {
		return err
	}
	log.Println("Created admin user")

	// CREATE DYNAMIC CONFIG
	err = repos.Config.Create(ctx, dynamicConfig)
	if err != nil {
		return err
	}
	log.Println("Created dynamic config")

	// CREATE POLICY
	err = repos.Policy.Create(ctx, policy.NewPolicy())
	if err != nil {
		return err
	}
	log.Println("Created policy")

	// CREATE REPORT
	err = repos.Report.Create(ctx, report.NewReport())
	if err != nil {
		return err
	}
	log.Println("Created report")

	return nil
}

func DefaultSeedConditional(staticConfig config.StaticConfig, dynamicConfig *config.DynamicConfig, repos *util.Store) error {
	if staticConfig.DB.Seed {
		return DefaultSeed(staticConfig, dynamicConfig, repos)
	}
	return nil
}
