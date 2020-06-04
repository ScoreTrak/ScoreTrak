package client

import (
	"ScoreTrak/pkg/check"
	"fmt"
)

type checkClient struct {
	s ScoretrakClient
}

func NewCheckClient(c ScoretrakClient) check.Serv {
	return &checkClient{c}
}

func (s checkClient) GetAllByRoundID(rID uint64) ([]*check.Check, error) {
	var chk []*check.Check
	err := genericGet(&chk, fmt.Sprintf("/check/%d", rID), s.s)
	if err != nil {
		return nil, err
	}
	return chk, nil
}

func (s checkClient) GetByRoundServiceID(rID uint64, sID uint64) ([]*check.Check, error) {
	var chk []*check.Check
	err := genericGet(&chk, fmt.Sprintf("/check/%d/%d", rID, sID), s.s)
	if err != nil {
		return nil, err
	}
	return chk, nil
}
