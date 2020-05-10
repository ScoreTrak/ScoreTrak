package check

// Check model contains an instance of a single check performed on a single host at a given round for a given service
type Check struct {
	Id int64 `json:"id,omitempty"`

	ServiceId int64 `json:"service_id,omitempty"`

	RoundId int64 `json:"round_id,omitempty"`

	// Represents an comment/log of a check. This will be helpful for debugging purposes during the competition
	Log string `json:"log,omitempty"`

	Passed bool `json:"passed,omitempty"`
}
