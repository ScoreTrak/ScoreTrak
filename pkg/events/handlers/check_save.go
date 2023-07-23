package handlers

import (
	"context"
	"encoding/json"
	"github.com/ScoreTrak/ScoreTrak/pkg/entities"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"time"
)

type CheckSavePayload struct {
	Check *entities.Check
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

func NewCheckSavePayload(check *entities.Check) *CheckSavePayload {
	return &CheckSavePayload{
		Check: check,
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

	_, err := CheckSavePayloadFromBytes(msg.Payload)
	if err != nil {
		return err
	}
	//
	//if err != nil {
	//	chk, err = s.entitiesClient.Check.Create().SetLog(logOutput).SetError(err.Error()).SetPassed(passed).SetRoundID(payload.RoundID).SetHostserviceID(payload.HostServiceID).Save(ctx)
	//	if err != nil {
	//		return err
	//	}
	//} else {
	//	chk, err = s.entitiesClient.Check.Create().SetLog(logOutput).SetPassed(passed).SetRoundID(payload.RoundID).SetHostserviceID(payload.HostServiceID).Save(ctx)
	//	if err != nil {
	//		return err
	//	}
	//}
	//
	//s.logger.Infof("Saved Check %s", chk.ID)
	msg.Ack()
	return nil
}
