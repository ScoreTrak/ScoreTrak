package outcome_writer

import (
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type OutcomeWriter struct {
	outcome *outcome.Outcome
	Logger  *otelzap.SugaredLogger
}

type OutcomeWriterOption func(ow *OutcomeWriter)

func NewOutcomeWriter(options ...OutcomeWriterOption) *OutcomeWriter {
	ow := &OutcomeWriter{}

	for _, o := range options {
		o(ow)
	}

	if ow.Logger == nil {
		l, _ := zap.NewProduction()
		ow.Logger = otelzap.New(l).Sugar()
	}

	return ow
}

func WithOutcome(outcome *outcome.Outcome) OutcomeWriterOption {
	return func(ow *OutcomeWriter) {
		ow.outcome = outcome
	}
}

func WithLogger(logger *otelzap.SugaredLogger) OutcomeWriterOption {
	return func(s *OutcomeWriter) {
		s.Logger = logger
	}
}

func (o *OutcomeWriter) SetStatus(status outcome.OUTCOME_STATUS) {
	o.outcome.Status = status
}

func (o *OutcomeWriter) SetError(e error) {
	o.outcome.Error = e
}
