package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/internal/entities/check"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type CheckSavePayload struct {
	Outcome *outcome.Outcome
	CheckID string
}

func (sp *CheckSavePayload) Bytes() ([]byte, error) {
	payloadBytes, err := json.Marshal(sp)
	if err != nil {
		return nil, err
	}
	return payloadBytes, nil
}

func CheckSavePayloadFromBytes(payloadBytes []byte) (*CheckSavePayload, error) {
	payload := CheckSavePayload{}
	err := json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func NewCheckSavePayload(outcome *outcome.Outcome, checkId string) *CheckSavePayload {
	return &CheckSavePayload{
		Outcome: outcome,
		CheckID: checkId,
	}
}

type CheckSaveHandler struct {
	entitiesClient *entities.Client
	logger         *otelzap.SugaredLogger
}

func NewCheckSaveHandler(entitiesClient *entities.Client, logger *otelzap.SugaredLogger) *CheckSaveHandler {
	return &CheckSaveHandler{
		entitiesClient: entitiesClient,
		logger:         logger,
	}
}

func (s *CheckSaveHandler) Handler(msg *message.Message) error {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*30))
	defer cancel()

	payload, err := CheckSavePayloadFromBytes(msg.Payload)
	if err != nil {
		return err
	}

	s.logger.Infow("Saving queue message", "check_id", payload.CheckID)
	chk, err := s.entitiesClient.Check.UpdateOneID(payload.CheckID).SetError(payload.Outcome.Error.Error()).SetOutcomeStatus(check.OutcomeStatus(payload.Outcome.Status)).Save(ctx)
	if err != nil {
		return err
	}

	s.logger.Infof("Updated Check %s", chk.ID)
	msg.Ack()
	return nil
}
