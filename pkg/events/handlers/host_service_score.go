package handlers

import (
	"context"
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec/resolver"
	"github.com/ScoreTrak/ScoreTrak/pkg/scorer/scorerservice"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"time"
)

type HostServiceScorePayload struct {
	RoundID       string
	RoundNumber   int
	HostAddress   string
	ServiceType   scorerservice.Service
	HostServiceID string
	Properties    []*entities.Property
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

func NewHostServiceScorePayload(roundID string, roundNumber int, hostAddress string, serviceType scorerservice.Service, hostServiceID string, properties []*entities.Property) *HostServiceScorePayload {
	return &HostServiceScorePayload{
		RoundID:       roundID,
		RoundNumber:   roundNumber,
		HostAddress:   hostAddress,
		ServiceType:   serviceType,
		HostServiceID: hostServiceID,
		Properties:    properties,
	}
}

type HostServiceScoreHandler struct {
	entitiesClient *entities.Client
	logger         *otelzap.SugaredLogger
}

func NewHostServiceScoreHandler(entitiesClient *entities.Client, logger *otelzap.SugaredLogger) *HostServiceScoreHandler {
	return &HostServiceScoreHandler{
		entitiesClient: entitiesClient,
		logger:         logger,
	}
}

func (s *HostServiceScoreHandler) Handler(msg *message.Message) error {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*30))
	defer cancel()

	payload, err := HostServiceScorePayloadFromBytes(msg.Payload)
	if err != nil {
		return err
	}

	executable := resolver.ExecutableByName(payload.ServiceType)

	properties := make(map[string]string)
	for _, v := range payload.Properties {
		properties[v.Key] = v.Value
	}

	err = exec.UpdateExecutableProperties(executable, properties)
	if err != nil {
		return err
	}

	e := exec.NewExec(ctx, payload.HostAddress, executable)

	passed, logOutput, err := e.Execute()

	var chk *entities.Check

	if err != nil {
		chk, err = s.entitiesClient.Check.Create().SetLog(logOutput).SetError(err.Error()).SetPassed(passed).SetRoundID(payload.RoundID).SetHostserviceID(payload.HostServiceID).Save(ctx)
		if err != nil {
			return err
		}
	} else {
		chk, err = s.entitiesClient.Check.Create().SetLog(logOutput).SetPassed(passed).SetRoundID(payload.RoundID).SetHostserviceID(payload.HostServiceID).Save(ctx)
		if err != nil {
			return err
		}
	}

	s.logger.Infof("Saved Check %s", chk.ID)
	msg.Ack()
	return nil
}
