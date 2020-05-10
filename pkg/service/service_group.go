package service

// Service Group model describes a groupping of services.
type ServiceGroup struct {
	Id int64 `json:"id,omitempty"`

	Name string `json:"name"`

	Label string `json:"label,omitempty"`

	// Enables or Disables the service
	Enabled bool `json:"enabled,omitempty"`
}
