package scorer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"sync"

	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

const (
	SCORER_NAME_REGEX = "^[a-z0-9-]$"
)

type Scorer struct {
	logger   *otelzap.SugaredLogger
	mu       sync.Mutex
	services map[ServiceType]ExecutorEntry
}

type ScorerOption func(scorer *Scorer)

func NewScorer(options ...ScorerOption) *Scorer {
	scorer := &Scorer{}
	for _, o := range options {
		o(scorer)
	}

	if scorer.logger == nil {
		l, _ := zap.NewProduction()
		scorer.logger = otelzap.New(l).Sugar()
	}
	return scorer
}

func WithLogger(logger *otelzap.SugaredLogger) ScorerOption {
	return func(s *Scorer) {
		s.logger = logger
	}
}

func (s *Scorer) Handle(name ServiceType, handler Executor) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if name == "" && regexp.MustCompile(SCORER_NAME_REGEX).MatchString(string(name)) {
		panic("scorer: invalid pattern")
	}
	if handler == nil {
		panic("scoreservice: nil scoreservice")
	}

	if _, exists := s.services[name]; exists {
		panic("scoreservice: multiple registrations for " + name)
	}

	if s.services == nil {
		s.services = make(map[ServiceType]ExecutorEntry)
	}

	e := ExecutorEntry{
		Name:     name,
		Executor: handler,
	}

	s.services[name] = e
}

var SCORER_HANDLER_NOT_FOUND = errors.New("no scorer handler found")

func (s *Scorer) Score(ctx context.Context, serviceType ServiceType, properties any) *outcome.Outcome {
	o := outcome.DefaultOutcome()
	ow := outcome_writer.NewOutcomeWriter(outcome_writer.WithLogger(s.logger), outcome_writer.WithOutcome(o))

	propertiesJson, err := json.Marshal(properties)
	if err != nil {
		ow.SetStatus(outcome.OUTCOME_STATUS_FAILED)
		ow.SetError(fmt.Errorf("unable able to convert properties into json: %v", err))
		return o
	}

	if service, exists := s.services[serviceType]; exists {
		service.Executor(ctx, ow, propertiesJson)
		return o
	} else {
		ow.SetStatus(outcome.OUTCOME_STATUS_FAILED)
		ow.SetError(SCORER_HANDLER_NOT_FOUND)
		return o
	}
}
