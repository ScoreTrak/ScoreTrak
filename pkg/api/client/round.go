package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/round"
)

type RoundClient struct {
	s ScoretrakClient
}

func NewRoundClient(c ScoretrakClient) round.Serv {
	return &RoundClient{c}
}

func (r RoundClient) GetLastNonElapsingRound() (*round.Round, error) {
	sg := &round.Round{}
	err := r.s.genericGet(sg, fmt.Sprintf("/round/last_non_elapsing"))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (r RoundClient) GetAll() ([]*round.Round, error) {
	var sg []*round.Round
	err := r.s.genericGet(&sg, fmt.Sprintf("/round"))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (r RoundClient) GetByID(id uint64) (*round.Round, error) {
	sg := &round.Round{}
	err := r.s.genericGet(sg, fmt.Sprintf("/round/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}
