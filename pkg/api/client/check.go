package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
)

type checkClient struct {
	s ScoretrakClient
}

func NewCheckClient(c ScoretrakClient) check.Serv {
	return &checkClient{c}
}

func (s checkClient) GetAllByRoundID(rID uint64) ([]*check.Check, error) {
	var chk []*check.Check
	err := s.s.genericGet(&chk, fmt.Sprintf("/check/%d", rID))
	if err != nil {
		return nil, err
	}
	return chk, nil
}

func (s checkClient) GetByRoundServiceID(rID uint64, sID uint64) (*check.Check, error) {
	var chk *check.Check
	err := s.s.genericGet(&chk, fmt.Sprintf("/check/%d/%d", rID, sID))
	if err != nil {
		return nil, err
	}
	return chk, nil
}
