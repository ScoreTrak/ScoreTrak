package grifts

import (
	"scoretrak/models"
	. "github.com/markbates/grift/grift"
)

var _ = Namespace("user", func() {
	Desc("admin", "Generates admin account with Username: admin, and Password: changeme")
	Add("admin", func(c *Context) error {
		u := &models.User{Username: "admin", Password: "changeme", PasswordConfirmation: "changeme"}
		return models.DB.Create(u)
	})

})
