package models

type Team struct {
	BaseModel
	NameModel
	RoundDuration                  string `json:"round_duration"`
	CurrentRoundId                 string `json:"current_round_id"`
	ViewableToPublic               string `json:"viewable_to_public"`
	IgnoreInCompleteRoundInScoring string `json:"ignore_in_complete_round_in_scoring"`
}
