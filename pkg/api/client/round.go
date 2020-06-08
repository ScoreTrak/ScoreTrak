package client

import (
	"ScoreTrak/pkg/round"
	"fmt"
)

type roundClient struct {
	s ScoretrakClient
}

func NewRoundClient(c ScoretrakClient) round.Serv {
	return &roundClient{c}
}

func (r roundClient) GetLastRound() (*round.Round, error) {
	sg := &round.Round{}
	err := genericGet(sg, fmt.Sprintf("/round"), r.s)
	if err != nil {
		return nil, err
	}
	return sg, nil
}
