package conv

import (
	"go.opentelemetry.io/otel/attribute"
)

const (
	ATTRIBUTE_PREFIX     = "scoretrak."
	COMPETITION_ID_KEY   = attribute.Key(ATTRIBUTE_PREFIX + "competition.id")
	COMPETITION_NAME_KEY = attribute.Key(ATTRIBUTE_PREFIX + "competition.name")
	TEAM_ID_KEY          = attribute.Key(ATTRIBUTE_PREFIX + "team.id")
	TEAM_NAME_KEY        = attribute.Key(ATTRIBUTE_PREFIX + "team.name")
)
