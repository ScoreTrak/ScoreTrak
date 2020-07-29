package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/check"
)

type CheckClient struct {
	s ScoretrakClient
}

func NewCheckClient(c ScoretrakClient) check.Serv {
	return &CheckClient{c}
}

func (s CheckClient) GetAllByRoundID(rID uint64) ([]*check.Check, error) {
	var chk []*check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/%d", rID))
	if err != nil {
		return nil, err
	}
	return chk, nil
}

func (s CheckClient) GetByRoundServiceID(rID uint64, sID uint64) (*check.Check, error) {
	var chk *check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/%d/%d", rID, sID))
	if err != nil {
		return nil, err
	}
	return chk, nil
}
