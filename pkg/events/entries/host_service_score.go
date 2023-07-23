package entries

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/events"
	"github.com/ScoreTrak/ScoreTrak/pkg/events/handlers"
	"github.com/ThreeDotsLabs/watermill/message"
)

const (
	HOST_SERVICE_SCORE_HANDLER_NAME = "host_service_score"
	TOPIC_HOST_SERVICE_SCORE        = "host_service_score"
)

func NewHostServiceScoreNoPublishHandlerEntry(sub message.Subscriber, hndlr *handlers.HostServiceScoreHandler) *events.NoPublishHandlerEntry {
	return events.NewNoPublishHandlerEntry(HOST_SERVICE_SCORE_HANDLER_NAME, TOPIC_HOST_SERVICE_SCORE, sub, hndlr.Handler)
}
