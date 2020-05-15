package property

// Property model describes a single key value pair for a service(parameters). An example could be a port for HTTP checking
type Property struct {
	ID uint64 `json:"id,omitempty"`

	ServiceID uint64 `json:"service_id" gorm:"not null"`

	Key string `json:"key,omitempty" gorm:"not null"`

	Value string `json:"value" gorm:"not null"`

	Status string `json:"status,omitempty" gorm:"not null default: View"`
}
