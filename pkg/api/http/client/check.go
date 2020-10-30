package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/gofrs/uuid"
)

type CheckClient struct {
	s ScoretrakClient
}

func NewCheckClient(c ScoretrakClient) *CheckClient {
	return &CheckClient{c}
}

func (s CheckClient) GetAllByRoundID(rID uint) ([]*check.Check, error) {
	var chk []*check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/round/%d", rID))
	if err != nil {
		return nil, err
	}
	return chk, nil
}

func (s CheckClient) GetAllByServiceID(sID uuid.UUID) ([]*check.Check, error) {
	var chk []*check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/service/%s", sID.String()))
	if err != nil {
		return nil, err
	}
	return chk, nil
}

func (s CheckClient) GetByRoundServiceID(rID uint, sID uuid.UUID) (*check.Check, error) {
	var chk *check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/round/%d/service/%s", rID, sID.String()))
	if err != nil {
		return nil, err
	}
	return chk, nil
}
