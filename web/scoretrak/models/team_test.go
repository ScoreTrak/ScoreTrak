package models

func (ms *ModelSuite) Test_Team_Create() {
	count, err := ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &Team{
		Name: "TestName",
		Role: "blue",
	}
	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	count, err = ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(1, count)
}

func (ms *ModelSuite) Test_Team_Create_ValidationErrors() {
	count, err := ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &Team{
		Name: "TestName",
		Role: "nonexistent",
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())

	count, err = ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)
}

func (ms *ModelSuite) Test_Team_Create_TeamExists() {
	count, err := ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(0, count)

	t := &Team{
		Name: "DuplicateName",
		Role: "black",
	}

	verrs, err := ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.False(verrs.HasAny())
	count, err = ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(1, count)

	t = &Team{
		Name: "DuplicateName",
		Role: "red",
	}

	verrs, err = ms.DB.ValidateAndCreate(t)
	ms.NoError(err)
	ms.True(verrs.HasAny())
	count, err = ms.DB.Count("teams")
	ms.NoError(err)
	ms.Equal(1, count)
}
