#package orm(pkg/storage/orm)

This package contains tests, and methods for creating, deleting, updating, retrieving models using https://github.com/go-gorm/gorm library.

Every file with {model_name}_repo.go file has a struct named {model_name}Repo that implements the respective interface located at pkg/{model_name}/{model_name}_repo/repo.go file.
These interfaces typically consist of the following methods:
 - Get: Gets the first row, this is a method that usually exists in models that only have one row
 - GetAll: Gets all entries
 - GetByID: Get a specific model by ID
 - Store: Stores a new model in the respective table
 - Upsert: Creates new entries, unless the entries already exist(this is useful when loading the competition. Check the details at pkg/competition/competitionservice/serv.go)
 - Update: Updates the entry
 - Delete: Deletes the entry
 - \<Other\>: Model specific methods.

All the methods take the context, that is then passed down to gorm when making the request for cancellation propagation.

Every test that is spawned locally will utilize pkg/storage/orm/dev-config.yml as input config by default.
