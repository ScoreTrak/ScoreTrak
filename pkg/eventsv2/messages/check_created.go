package messages

import (
	"github.com/scoretrak/scoretrak/pkg/scorer"
)

type ChecksCreatedMessage struct {
	RoundID       string             `json:"round_id"`
	RoundNumber   int                `json:"round_number"`
	TeamID        string             `json:"team_id"`
	HostAddress   string             `json:"host_address"`
	CheckID       string             `json:"check_id"`
	ServiceType   scorer.ServiceType `json:"service_type"`
	HostServiceID string             `json:"host_service_id"`
	Properties    map[string]string  `json:"properties"`
}
