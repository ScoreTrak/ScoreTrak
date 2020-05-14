package service_group

import "ScoreTrak/pkg/service"

// Serv Group model describes a grouping of services.
type ServiceGroup struct {
	ID uint64 `json:"id,omitempty"`

	Name string `json:"name"`

	Services []service.Service `gorm:"foreignkey:ServiceGroupID"`

	// Enables or Disables the service
	Enabled bool `json:"enabled,omitempty"`
}

func (ServiceGroup) TableName() string {
	return "service_groups"
}
