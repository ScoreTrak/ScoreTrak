package config

type Repo interface {
	Get() (*DynamicConfig, error)
	Update(*DynamicConfig) error
	ResetScores() error
	DeleteCompetition() error
}
