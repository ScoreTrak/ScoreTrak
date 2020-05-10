package team

// Team model represents internal team model of the scoring engine.
type Team struct {

	// this id refers to ID of a team in web.
	Id int64 `json:"id"`

	Enabled bool `json:"enabled,omitempty"`
}
