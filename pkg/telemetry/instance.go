package telemetry

import "github.com/oklog/ulid/v2"

type Instance struct {
	ID string
}

func (i *Instance) String() string {
	return i.ID
}

func NewInstance() *Instance {
	return &Instance{ID: ulid.Make().String()}
}
