package client

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/competition"
)

type CompetitionClient struct {
	s ScoretrakClient
}

func NewCompetitionClient(c ScoretrakClient) *CompetitionClient {
	return &CompetitionClient{c}
}

func (s CompetitionClient) FetchCoreCompetition() (*competition.Competition, error) {
	var sg *competition.Competition
	err := s.s.GenericGet(&sg, "/competition/export_core")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s CompetitionClient) FetchEntireCompetition() (*competition.Competition, error) {
	var sg *competition.Competition
	err := s.s.GenericGet(&sg, "/competition/export_all")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s CompetitionClient) LoadCompetition(u *competition.Competition) error {
	return s.s.GenericStore(u, "/competition/upload")
}
