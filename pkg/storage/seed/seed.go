package seed

import (
	"context"
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/storage/util"
	"github.com/ScoreTrak/ScoreTrak/pkg/team"
	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgconn"
	"golang.org/x/crypto/bcrypt"
)

// create black team
// create admin user

func SeedDB(store *util.Store, staticConfig config.StaticConfig) error {
	ctx := context.Background()

	var firstUUID = uuid.FromStringOrNil("00000000-0000-0000-0000-000000000001")

	fls := false
	idx := uint64(0)
	err := store.Team.Store(ctx, []*team.Team{{
		ID:    firstUUID,
		Name:  "Black Team",
		Pause: &fls,
		Hide:  &fls,
		Index: &idx,
		Users: nil,
		Hosts: nil,
	}})
	if err != nil {
		var serr *pgconn.PgError
		ok := errors.As(err, &serr)
		if !ok || serr.Code != "23505" {
			return err
		}
	}

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(staticConfig.AdminPassword), bcrypt.DefaultCost)
	err = store.Users.Store(context.Background(), []*user.User{{
		ID:           firstUUID,
		Role:         user.Black,
		TeamID:       firstUUID,
		Username:     staticConfig.AdminUsername,
		PasswordHash: string(passwordHash),
	}})
	if err != nil {
		var serr *pgconn.PgError
		ok := errors.As(err, &serr)
		if !ok || serr.Code != "23505" {
			return err
		}
	}

	return nil
}
