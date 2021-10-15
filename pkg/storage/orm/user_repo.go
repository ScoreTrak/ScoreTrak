package orm

import (
	"context"
	"fmt"

	"github.com/ScoreTrak/ScoreTrak/pkg/user"
	"github.com/ScoreTrak/ScoreTrak/pkg/user/userrepo"
	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) userrepo.Repo {
	return &userRepo{db}
}

func (h *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	result := h.db.WithContext(ctx).Delete(&user.User{}, "id = ?", id)
	if result.Error != nil {
		return fmt.Errorf("error deleting user with id: %s, err: %w", id.String(), result.Error)
	}
	return nil
}

func (h *userRepo) GetAll(ctx context.Context) ([]*user.User, error) {
	users := make([]*user.User, 0)
	err := h.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (h *userRepo) GetByID(ctx context.Context, id uuid.UUID) (*user.User, error) {
	usr := &user.User{}
	err := h.db.WithContext(ctx).Where("id = ?", id).First(usr).Error
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (h *userRepo) GetByUsername(ctx context.Context, username string) (*user.User, error) {
	usr := &user.User{}
	err := h.db.WithContext(ctx).Where("username = ?", username).First(usr).Error
	if err != nil {
		return nil, err
	}
	return usr, nil
}

func (h *userRepo) Store(ctx context.Context, usr []*user.User) error {
	err := h.db.WithContext(ctx).Create(usr).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *userRepo) Upsert(ctx context.Context, usr []*user.User) error {
	err := h.db.WithContext(ctx).Clauses(clause.OnConflict{DoNothing: true}).Create(usr).Error
	if err != nil {
		return err
	}
	return nil
}

func (h *userRepo) Update(ctx context.Context, usr *user.User) error {
	err := h.db.WithContext(ctx).Model(usr).Updates(user.User{PasswordHash: usr.PasswordHash, Username: usr.Username, TeamID: usr.TeamID, Role: usr.Role}).Error
	if err != nil {
		return err
	}
	return nil
}
