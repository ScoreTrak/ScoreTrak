package handlers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/scoretrak/scoretrak/internal/entities"
	"github.com/scoretrak/scoretrak/pkg/scorer"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
)

type HostServiceScorePayload struct {
	RoundID       string
	RoundNumber   int
	HostAddress   string
	CheckID       string
	ServiceType   scorer.ServiceType
	HostServiceID string
	Properties    map[string]string
}

func (sp *HostServiceScorePayload) Bytes() ([]byte, error) {
	payloadBytes, err := json.Marshal(sp)
	if err != nil {
		return nil, err
	}
	return payloadBytes, nil
}

func HostServiceScorePayloadFromBytes(payloadBytes []byte) (*HostServiceScorePayload, error) {
	payload := HostServiceScorePayload{}
	err := json.Unmarshal(payloadBytes, &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func NewHostServiceScorePayload(roundID string, roundNumber int, hostAddress string, serviceType scorer.ServiceType, hostServiceID string, checkID string, properties map[string]string) *HostServiceScorePayload {
	return &HostServiceScorePayload{
		RoundID:       roundID,
		RoundNumber:   roundNumber,
		HostAddress:   hostAddress,
		ServiceType:   serviceType,
		CheckID:       checkID,
		HostServiceID: hostServiceID,
		Properties:    properties,
	}
}

type HostServiceScoreHandler struct {
	entitiesClient *entities.Client
	// fb             *flagbearer.FlagBearer
	scorer *scorer.Scorer
	logger *otelzap.SugaredLogger
}

func NewHostServiceScoreHandler(entitiesClient *entities.Client, scorer *scorer.Scorer, logger *otelzap.SugaredLogger) *HostServiceScoreHandler {
	return &HostServiceScoreHandler{
		entitiesClient: entitiesClient,
		scorer:         scorer,
		logger:         logger,
	}
}

func (s *HostServiceScoreHandler) Handler(msg *message.Message) ([]*message.Message, error) {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*30))
	defer cancel()

	s.logger.Infoln(msg)

	payload, err := HostServiceScorePayloadFromBytes(msg.Payload)
	if err != nil {
		return nil, err
	}

	s.logger.Infoln(payload)
	s.logger.Infow("Scoring host service", "host_service_id", payload.HostServiceID)
	outcome := s.scorer.Score(ctx, payload.ServiceType, payload.Properties)
	if outcome.Error != nil {
		return nil, err
	}

	var msgs []*message.Message
	// Return message back to queue
	checkSavePayload := NewCheckSavePayload(outcome, payload.CheckID)
	checkSavePayloadBytes, err := checkSavePayload.Bytes()
	if err != nil {
		return nil, err
	}
	checkSaveMsg := message.NewMessage(watermill.NewULID(), checkSavePayloadBytes)
	msgs = append(msgs, checkSaveMsg)

	return msgs, nil
}
