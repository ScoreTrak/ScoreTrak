package swarm

type Swarm struct {
	Id int64 `json:"id,omitempty"`

	ServiceId int64 `json:"service_id"`

	Label string `json:"label"`
}
