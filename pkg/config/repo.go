package config

type Repo interface {
	Get() (*Config, error)
	Update(*Config) error
}
