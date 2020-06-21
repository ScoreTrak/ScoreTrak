package grifts

import (
	. "github.com/markbates/grift/grift"
	"scoretrak/constants"
	"scoretrak/models"
)

var _ = Namespace("create", func() {
	Desc("admin", "Generates admin account and group. Username: admin, Password: changeme")
	Add("admin", func(c *Context) error {
		t := &models.Team{Name: "Black Team", Role: constants.Black}
		_, err := models.DB.ValidateAndCreate(t)
		if err != nil {
			return err
		}

		b_t := []models.Team{}
		query := models.DB.Where("name = 'Black Team'")
		err = query.All(&b_t)
		if err != nil {
			return err
		}

		u := &models.User{Username: "admin", Password: "changeme", PasswordConfirmation: "changeme", TeamID: b_t[0].ID}
		_, err = u.Create(models.DB)
		if err != nil {
			return err
		}

		a_t := []models.User{}
		query = models.DB.Where("username = 'admin'")
		err = query.All(&a_t)
		if err != nil {
			return err
		}
		return nil
	})
})
