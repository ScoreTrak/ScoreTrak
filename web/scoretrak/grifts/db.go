package grifts

import (
	. "github.com/markbates/grift/grift"
	"scoretrak/models"
	
)

var _ = Namespace("user", func() {
	Desc("admin", "Generates admin account with Username: admin, and Password: changeme")
	Add("admin", func(c *Context) error {
		u := &models.User{Username: "admin", Password: "changeme", PasswordConfirmation: "changeme"}
		_, err := u.Create(models.DB)
		if err != nil {
			return err
		}
		return nil
	})
})
