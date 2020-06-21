package models

import (
	"github.com/gofrs/uuid"
	"log"
)

func (ms *ModelSuite) Test_User_Create() {
	defer destroy_Users_And_Teams(ms)
	ms.LoadFixture("Create a Team")
	t, terr := GetTeamByName(ms.DB, "Black Team")
	if terr != nil {
		log.Fatal(terr)
	}
	u := &User{
		Username:             "testuser",
		Password:             "testpass",
		PasswordConfirmation: "testpass",
		TeamID:               t.ID,
	}

	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {
	defer destroy_Users_And_Teams(ms)
	u := &User{
		Password: "password",
	}

	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_User_Create_UserExists() {
	defer destroy_Users_And_Teams(ms)
	ms.LoadFixture("Create a Team")
	t, terr := GetTeamByName(ms.DB, "Black Team")
	if terr != nil {
		log.Fatal(terr)
	}
	u := &User{
		Username:             "testuser",
		Password:             "testpass",
		PasswordConfirmation: "testpass",
		TeamID:               t.ID,
	}

	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	u = &User{
		Username:             "testuser",
		Password:             "testpass",
		PasswordConfirmation: "testpass",
		TeamID:               t.ID,
	}

	verrs, err = u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_User_Create_NonExistentTeamID() {
	defer destroy_Users_And_Teams(ms)
	s := "6ba7b810-9dad-11d1-80b4-00c04fd430c8" //Non-existent UUID.
	u3, err := uuid.FromString(s)
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", s, err)
	}

	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	u := &User{
		Username:             "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
		TeamID:               u3,
	}

	ms.Zero(u.PasswordHash)
	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}

func (ms *ModelSuite) Test_User_Update_LastBlackTeamUser() {
	defer destroy_Users_And_Teams(ms)
	ms.LoadFixture("Create Multiple Teams")
	t, terr := GetTeamByName(ms.DB, "Black Team")
	if terr != nil {
		log.Fatal(terr)
	}
	u := &User{
		Username:             "testuser",
		Password:             "testpass",
		PasswordConfirmation: "testpass",
		TeamID:               t.ID,
	}

	count, err := ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(0, count)

	ms.Zero(u.PasswordHash)

	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	ms.NotZero(u.PasswordHash)

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)

	t, terr = GetTeamByName(ms.DB, "Red Team")
	if terr != nil {
		log.Fatal(terr)
	}

	t_u, err := GetUserByUsername(ms.DB, "testuser")
	if err != nil {
		log.Fatal(err)
	}
	t_u.Username = "testusername2"
	verrs, err = ms.DB.ValidateAndUpdate(&t_u)
	ms.NoError(err)
	ms.False(verrs.HasAny())

	t_u.TeamID = t.ID
	verrs, err = ms.DB.ValidateAndUpdate(&t_u)
	ms.NoError(err)
	ms.True(verrs.HasAny())

}

func destroy_Users_And_Teams(ms *ModelSuite) {
	users := []User{}
	err := ms.DB.All(&users)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(users); i++ {
		user := users[i]
		ms.DB.Destroy(&user)
	}

	teams := []Team{}
	err = ms.DB.All(&teams)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(teams); i++ {
		team := teams[i]
		ms.DB.Destroy(&team)
	}
}
