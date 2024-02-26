package models

type Competition struct {
	BaseModel
	NameModel
	RoundDuration                  string `json:"round_duration"`
	CurrentRoundId                 string `json:"current_round_id"`
	ViewableToPublic               string `json:"viewable_to_public"`
	IgnoreInCompleteRoundInScoring string `json:"ignore_in_complete_round_in_scoring"`
	//field.String("name").Match(regexp.MustCompile("^[a-z0-9_]*$")).MinLen(4).MaxLen(32).Unique(),
	//field.String("display_name").Match(regexp.MustCompile("^[a-zA-Z0-9\\s]*$")).MinLen(4).MaxLen(64),
	//field.Int("round_duration").Optional().Default(60),
	//field.String("current_round_id").Optional().Nillable().Comment("Most recently completed round"),
	//field.Bool("viewable_to_public").Nillable().Optional(),
	//field.Bool("ignore_incomplete_round_in_scoring").Optional(),
	//field.Time("to_be_started_at").Nillable().Optional(),
	//field.Time("started_at").Nillable().Optional(),
	//field.Time("finished_at").Nillable().Optional(),
}
