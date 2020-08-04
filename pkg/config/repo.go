package config

type Repo interface {
	Get() (*DynamicConfig, error)
	Update(*DynamicConfig) error
	ResetCompetition() error
	DeleteCompetition() error
}
