package service

// Service Model represents a service that is being scored for a given host
type Service struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	// Points granted for a successful check
	Points int64 `json:"points"`

	// The frequency of a service check. If round_units is 5 and round_delay is 0, then service checks will happen on every 5th round. (5,10, etc)
	RoundUnits int64 `json:"round_units,omitempty"`

	// The frequency of a service check. If round_units is 7 and round_delay is 3, then service checks will happen on every 7th round with an offset of 3. (10,17, etc)
	RoundDelay int64 `json:"round_delay,omitempty"`

	// ID of a service group the service belongs to
	ServiceGroupId int64 `json:"service_group_id"`

	// ID of a host the service belongs to
	HostId int64 `json:"host_id"`

	// Enables or Disables the service
	Enabled bool `json:"enabled,omitempty"`
}
