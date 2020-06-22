package docker

import "github.com/L1ghtman2k/ScoreTrak/pkg/config"

type Docker struct {
	IsSwarm bool
}

func NewDocker() *Docker {
	return &Docker{}
}

func (Docker) DeployServiceGroup(label string, config *config.StaticConfig) error {

}

func (Docker) GetInventory() (error, SimpleCluster) {
	return nil, SimpleCluster{}
}

type SimpleCluster struct {
	Hosts []SimpleHost
}

type SimpleHost struct {
	IsManager bool
	Labels    []string
}
