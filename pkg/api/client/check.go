package client

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/check"
	"github.com/gofrs/uuid"
)

type CheckClient struct {
	s ScoretrakClient
}

func NewCheckClient(c ScoretrakClient) check.Serv {
	return &CheckClient{c}
}

func (s CheckClient) GetAllByRoundID(rID uint) ([]*check.Check, error) {
	var chk []*check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/%d", rID))
	if err != nil {
		return nil, err
	}
	return chk, nil
}

func (s CheckClient) GetByRoundServiceID(rID uint, sID uuid.UUID) (*check.Check, error) {
	var chk *check.Check
	err := s.s.GenericGet(&chk, fmt.Sprintf("/check/%d/%s", rID, sID.String()))
	if err != nil {
		return nil, err
	}
	return chk, nil
}
