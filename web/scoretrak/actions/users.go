package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/x/responder"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"
	"net/http"
	"scoretrak/constants"
	"scoretrak/models"
)

func UsersList(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	users := []models.User{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Users from the DB
	if err := q.All(&users); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// Add the paginator to the context so it can be used in the template.
		c.Set("pagination", q.Paginator)
		c.Set("users", users)

		return c.Render(http.StatusOK, r.HTML("/users/index.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(users))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(users))
	}).Respond(c)

}

func UsersShow(c buffalo.Context) error {

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		c.Set("user", user)

		return c.Render(http.StatusOK, r.HTML("/users/show.plush.html"))
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(200, r.JSON(user))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(200, r.XML(user))
	}).Respond(c)
}

func UsersEdit(c buffalo.Context) error {
	// Get the DB connection from the context

	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}
	c.Set("user", user)
	teams := []models.Team{}
	if err := tx.All(&teams); err != nil {
		return err
	}
	m := teamToIDMap(teams)
	c.Set("teams", m)

	return c.Render(http.StatusOK, r.HTML("/users/edit.plush.html"))
}

func UsersUpdate(c buffalo.Context) error {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	// Bind Team to the html form elements
	if err := c.Bind(user); err != nil {
		return err
	}

	verrs, err := user.Update(tx)
	if err != nil {
		return err
	}

	if verrs.HasAny() {
		return responder.Wants("html", func(c buffalo.Context) error {
			// Make the errors available inside the html template
			c.Set("errors", verrs)

			// Render again the edit.html template that the user can
			// correct the input.
			c.Set("user", user)
			return c.Render(http.StatusUnprocessableEntity, r.HTML("/users/edit.plush.html"))
		}).Wants("json", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.JSON(verrs))
		}).Wants("xml", func(c buffalo.Context) error {
			return c.Render(http.StatusUnprocessableEntity, r.XML(verrs))
		}).Respond(c)
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a success message
		c.Flash().Add("success", T.Translate(c, "user.updated.success"))

		// and redirect to the show page
		return c.Redirect(http.StatusSeeOther, "/users/%v", user.ID)
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(user))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(user))
	}).Respond(c)
}

func UsersDestroy(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return fmt.Errorf("no transaction found")
	}

	// Allocate an empty User
	user := &models.User{}

	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(http.StatusNotFound, err)
	}

	if err := tx.Destroy(user); err != nil {
		return err
	}

	return responder.Wants("html", func(c buffalo.Context) error {
		// If there are no errors set a flash message
		c.Flash().Add("success", T.Translate(c, "user.destroyed.success"))

		// Redirect to the index page
		return c.Redirect(http.StatusSeeOther, "/users")
	}).Wants("json", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.JSON(user))
	}).Wants("xml", func(c buffalo.Context) error {
		return c.Render(http.StatusOK, r.XML(user))
	}).Respond(c)

}

//UsersNew renders the users form
func UsersNew(c buffalo.Context) error {
	u := models.User{}
	c.Set("user", u)

	tx := c.Value("tx").(*pop.Connection)
	teams := []models.Team{}
	if err := tx.All(&teams); err != nil {
		return err
	}
	m := teamToIDMap(teams)
	c.Set("teams", m)
	return c.Render(200, r.HTML("users/new.plush.html"))
}

// UsersCreate registers a new user with the application.
func UsersCreate(c buffalo.Context) error {
	u := &models.User{}

	if err := c.Bind(u); err != nil {
		return errors.WithStack(err)
	}

	tx := c.Value("tx").(*pop.Connection)
	verrs, err := u.Create(tx)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		c.Set("user", u)
		c.Set("errors", verrs)

		tx := c.Value("tx").(*pop.Connection)
		teams := []models.Team{}
		if err := tx.All(&teams); err != nil {
			return err
		}
		m := teamToIDMap(teams)
		c.Set("teams", m)

		return c.Render(200, r.HTML("users/new.plush.html"))
	}

	// c.Session().Set("current_user_id", u.ID)
	c.Flash().Add("success", "User Created!")

	return c.Redirect(302, "/users")
}

func teamToIDMap(teams []models.Team) map[string]uuid.UUID {
	m := make(map[string]uuid.UUID)

	for _, team := range teams {
		m[team.Name] = team.ID
	}

	return m
}

// SetCurrentUser attempts to find a user based on the current_user_id
// in the session. If one is found it is set on the context.
func SetCurrentUser(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid != nil {
			u := &models.User{}
			tx := c.Value("tx").(*pop.Connection)
			err := tx.Find(u, uid)
			if err != nil {
				c.Session().Clear()
				return c.Redirect(302, "/")
				// return errors.WithStack(err)
			}
			c.Set("current_user", u)
			t := &models.Team{}
			if u != nil {
				tx := c.Value("tx").(*pop.Connection)
				err := tx.Find(t, u.TeamID)
				if err != nil {
					return errors.WithStack(err)
				}
			}
			c.Set("current_team", t)
			c.Set("roles", constants.Roles)
		}
		return next(c)
	}
}

// Authorize require a user be logged in before accessing a route
func Authorize(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		if uid := c.Session().Get("current_user_id"); uid == nil {
			c.Session().Set("redirectURL", c.Request().URL.String())

			err := c.Session().Save()
			if err != nil {
				return errors.WithStack(err)
			}

			c.Flash().Add("danger", "You must be authorized to see that page")
			return c.Redirect(302, "/auth/new")
		}
		return next(c)
	}
}

//VerifyAdminTeam verifies weather or not a given user set by context in SetCurrentUser belongs to an Admin Team
func AuthorizeBlackTeam(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		r := c.Value("current_team").(*models.Team)
		if r.Role != constants.Black {
			c.Flash().Add("danger", fmt.Sprintf("You must be a member of a team with role \"%s\" to access this page", constants.Black))
			return c.Redirect(302, "/")
		}
		return next(c)
	}
}
