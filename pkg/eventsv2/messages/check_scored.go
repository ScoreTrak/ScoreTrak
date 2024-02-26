package messages

import (
	"encoding/json"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
)

type ChecksScoredMessage struct {
	Outcome *outcome.Outcome `json:"outcome"`
	CheckID string           `json:"check_id"`
}

func (csm *ChecksScoredMessage) UnmarshalJSON(data []byte) error {
	type Alias ChecksScoredMessage // Define an alias to avoid recursion

	aux := &struct {
		Outcome map[string]interface{} `json:"outcome"`
		CheckID string                 `json:"check_id"`
		*Alias
	}{
		Alias: (*Alias)(csm),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Unmarshal the outcome field
	if aux.Outcome != nil {
		outcomeData, err := json.Marshal(aux.Outcome)
		if err != nil {
			return err
		}
		var out outcome.Outcome
		if err := json.Unmarshal(outcomeData, &out); err != nil {
			return err
		}
		csm.Outcome = &out
	}

	csm.CheckID = aux.CheckID

	return nil
}
