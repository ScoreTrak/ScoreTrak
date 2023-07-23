package scorer

import "github.com/ScoreTrak/ScoreTrak/pkg/scorer/scorerservice"

type Scorer struct {
}

func (s *Scorer) Score(hostAddress string, serviceType scorerservice.Service, properties map[string]any) {

}

func (s *Scorer) ScoreAndSave() {

}

func (s *Scorer) checkIfPreviousRoundHasCompleted() {

}
