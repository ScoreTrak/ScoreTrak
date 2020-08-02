package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/round"
)

type RoundClient struct {
	s ScoretrakClient
}

func NewRoundClient(c ScoretrakClient) round.Serv {
	return &RoundClient{c}
}

func (r *RoundClient) GetLastNonElapsingRound() (*round.Round, error) {
	sg := &round.Round{}
	err := r.s.GenericGet(sg, fmt.Sprintf("/round/last_non_elapsing"))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (r *RoundClient) GetAll() ([]*round.Round, error) {
	var sg []*round.Round
	err := r.s.GenericGet(&sg, fmt.Sprintf("/round"))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (r *RoundClient) GetByID(id uint) (*round.Round, error) {
	sg := &round.Round{}
	err := r.s.GenericGet(sg, fmt.Sprintf("/round/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (r *RoundClient) GetLastRound() (*round.Round, error) {
	sg := &round.Round{}
	err := r.s.GenericGet(sg, fmt.Sprintf("/round/last"))
	if err != nil {
		return nil, err
	}
	return sg, nil
}
