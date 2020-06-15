package service_group

import "ScoreTrak/pkg/service"

// Serv Group model describes a grouping of services.
type ServiceGroup struct {
	ID uint64 `json:"id,omitempty"`

	Name string `json:"name" gorm:"not null;unique;default:null"`

	// Enables or Disables the service
	Enabled *bool `json:"enabled,omitempty" gorm:"not null;default: false"`

	Services []*service.Service `gorm:"foreignkey:ServiceGroupID"`
}

func (ServiceGroup) TableName() string {
	return "service_groups"
}
