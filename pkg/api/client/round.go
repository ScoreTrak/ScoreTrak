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

func (r roundClient) GetLastNonElapsingRound() (*round.Round, error) {
	sg := &round.Round{}
	err := r.s.genericGet(sg, fmt.Sprintf("/round"))
	if err != nil {
		return nil, err
	}
	return sg, nil
}
