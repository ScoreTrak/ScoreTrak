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
	err := repos.Team.Upsert(ctx, []*team.Team{{
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
	log.Println("Create Black Team")

	// CREATE ADMIN USER
	var adminUUID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(staticConfig.AdminPassword), bcrypt.DefaultCost)
	err = repos.Users.Upsert(context.Background(), []*user.User{{
		ID:           adminUUID,
		Role:         user.Black,
		TeamID:       blackTeamUUID,
		Username:     staticConfig.AdminUsername,
		PasswordHash: string(passwordHash),
	}})
	if err != nil {
		return err
	}

	// CREATE DYNAMIC CONFIG
	err = repos.Config.Upsert(ctx, dynamicConfig)
	if err != nil {
		return err
	}

	// CREATE POLICY
	err = repos.Policy.Upsert(ctx, policy.NewPolicy())
	if err != nil {
		return err
	}

	// CREATE REPORT
	err = repos.Report.Upsert(ctx, report.NewReport())
	if err != nil {
		return err
	}

	return nil
}
