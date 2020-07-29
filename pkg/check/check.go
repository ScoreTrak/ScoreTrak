package check

// Check model contains an instance of a single check performed on a single host at a given round for a given service
type Check struct {
	ServiceID uint32 `json:"service_id,omitempty" gorm:"primary_key;auto_increment:false"`

	RoundID uint32 `json:"round_id,omitempty" gorm:"primary_key;auto_increment:false"`

	// Represents an comment/log of a check. This will be helpful for debugging purposes during the competition
	Log string `json:"log,omitempty"`

	Err string `json:"err,omitempty"`

	Passed *bool `json:"passed,omitempty" gorm:"not null; default: false"`
}
