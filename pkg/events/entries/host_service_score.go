package entries

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/scoretrak/scoretrak/pkg/events"
	"github.com/scoretrak/scoretrak/pkg/events/handlers"
)

const (
	HOST_SERVICE_SCORE_HANDLER_NAME string = "host_service_score"
	TOPIC_HOST_SERVICE_SCORE        string = "host_service_score"
)

// func NewHostServiceScoreNoPublishHandlerEntry(sub message.Subscriber, hndlr *handlers.HostServiceScoreHandler) *events.NoPublishHandlerEntry {
// 	return events.NewNoPublishHandlerEntry(HOST_SERVICE_SCORE_HANDLER_NAME, TOPIC_HOST_SERVICE_SCORE, sub, hndlr.Handler)
// }

func NewHostServiceScoreHandlerEntry(pub message.Publisher, sub message.Subscriber, hndlr *handlers.HostServiceScoreHandler) *events.HandlerEntry {
	return events.NewHandlerEntry(HOST_SERVICE_SCORE_HANDLER_NAME, TOPIC_HOST_SERVICE_SCORE, sub, TOPIC_CHECK_SAVE, pub, hndlr.Handler)
}
