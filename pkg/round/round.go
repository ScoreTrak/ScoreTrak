package round

import (
	"time"
)

// Round Model holds the information about all the rounds that have passed
type Round struct {

	// Round ID can also represent the round number
	Id int64 `json:"id,omitempty"`

	Start time.Time `json:"start,omitempty"`

	End time.Time `json:"end,omitempty"`

	Passed bool `json:"passed,omitempty"`
}
