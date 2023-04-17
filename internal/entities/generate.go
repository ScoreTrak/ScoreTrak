package entities

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature privacy,schema/snapshot,sql/upsert,entql ./schema
//go:generate go run -mod=mod entc.go
