package entries

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/scoretrak/scoretrak/pkg/events"
	"github.com/scoretrak/scoretrak/pkg/events/handlers"
)

const (
	CHECK_SAVE_HANDLER_NAME string = "check_save"
	TOPIC_CHECK_SAVE        string = "check_save"
)

func NewCheckSaveNoPublishHandlerEntry(sub message.Subscriber, hndlr *handlers.CheckSaveHandler) *events.NoPublishHandlerEntry {
	return events.NewNoPublishHandlerEntry(CHECK_SAVE_HANDLER_NAME, TOPIC_CHECK_SAVE, sub, hndlr.Handler)
}
