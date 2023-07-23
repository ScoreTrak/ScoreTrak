package flagbearer

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities"
)

type FlagBearer struct {
	entitiesClient *entities.Client
}

func NewFlagBearer(cfg *config.Config, entitiesClient *entities.Client) *FlagBearer {
	return &FlagBearer{entitiesClient: entitiesClient}
}

func (f *FlagBearer) EndRound() {

}

func (f *FlagBearer) StartRound() {
	// TODO: Move logic on starting, running round to this component
}
