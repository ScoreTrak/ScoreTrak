package swarm

import "ScoreTrak/pkg/service"

type Swarm struct {
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	Service service.Service `json:"-" gorm:"foreignkey:ServiceID"`

	ServiceID *uint64 `json:"service_id" gorm:"not null"`

	Label string `json:"label" gorm:"not null"`
}

func (Swarm) TableName() string {
	return "swarms"
}
