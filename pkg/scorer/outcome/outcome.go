package outcome

import (
	"encoding/json"
	"fmt"
	"github.com/creasty/defaults"
)

const (
	// OUTCOME_STATUS_FAILED the scoring has passed.
	OUTCOME_STATUS_PASSED OUTCOME_STATUS = "PASSED"

	// OUTCOME_STATUS_FAILED the scoring has failed.
	OUTCOME_STATUS_FAILED OUTCOME_STATUS = "FAILED"

	// OUTCOME_STATUS_INVALID incorrect parameters have been set
	OUTCOME_STATUS_INVALID OUTCOME_STATUS = "INVALID"
)

// OUTCOME_STATUS status of the scoring outcome
type OUTCOME_STATUS string

type Outcome struct {
	Status OUTCOME_STATUS `json:"status" default:"FAILED"`
	Error  error          `json:"error"`
}

func (o *Outcome) UnmarshalJSON(data []byte) error {
	type Alias Outcome
	aux := &struct {
		Status string `json:"status"`
		Error  string `json:"error"`
		*Alias
	}{
		Alias: (*Alias)(o),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch aux.Status {
	case string(OUTCOME_STATUS_PASSED):
		o.Status = OUTCOME_STATUS_PASSED
	case string(OUTCOME_STATUS_FAILED):
		o.Status = OUTCOME_STATUS_FAILED
	case string(OUTCOME_STATUS_INVALID):
		o.Status = OUTCOME_STATUS_INVALID
	}

	if aux.Error != "" {
		o.Error = fmt.Errorf(aux.Error)
	} else {
		o.Error = nil
	}

	return nil
}

func DefaultOutcome() *Outcome {
	outcome := &Outcome{}
	if err := defaults.Set(outcome); err != nil {
		return nil
	}
	return outcome
}
