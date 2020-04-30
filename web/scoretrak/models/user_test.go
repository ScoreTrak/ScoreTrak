package models

import (
	"log"
	"github.com/gofrs/uuid"
)

func (ms *ModelSuite) Test_User_Create() {
	
	ms.LoadFixture("Create a Team")
	t, terr := get_Team(ms)
	if terr != nil{
		log.Fatal(terr)
	}
	u := &User{
		Username:             "testuser",
		Password:             "testpass",
		PasswordConfirmation: "testpass",
		TeamID:				  t.ID,
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

	destroy_Users(ms)
}

func (ms *ModelSuite) Test_User_Create_ValidationErrors() {

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
	ms.LoadFixture("Create a Team")
	t, terr := get_Team(ms)
	if terr != nil{
		log.Fatal(terr)
	}
	u := &User{
		Username:             "testuser",
		Password:             "testpass",
		PasswordConfirmation: "testpass",
		TeamID:				  t.ID,
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
		TeamID:				  t.ID,
	}

	verrs, err = u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("users")
	ms.NoError(err)
	ms.Equal(1, count)
	destroy_Users(ms)
}


func (ms *ModelSuite) Test_User_Create_NonExistentTeamID() {

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
		TeamID:				  u3,
	}

	ms.Zero(u.PasswordHash)
	verrs, err := u.Create(ms.DB)
	ms.NoError(err)
	ms.True(verrs.HasAny())
}


func get_Team(ms *ModelSuite) (Team, error){
	b_t := []Team{}
	query := ms.DB.Where("name = 'Mock Team'")
	err := query.All(&b_t)
	if err != nil {
		return Team{}, err
	}

	return b_t[0], err
}

func destroy_Users(ms *ModelSuite){
	users := []User{}
	err := ms.DB.All(&users)
	if err != nil{
		panic(err)
	}
	for i := 0; i < len(users); i++ {
		user := users[i]
		ms.DB.Destroy(&user)
	}
}

