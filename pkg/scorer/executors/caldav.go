package executors

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/emersion/go-webdav"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"net/http"
	"strings"
)

// TODO(thisisibrahimd) validate properties with golang-validator
type CalDavProperties struct {
	Scheme         string `json:"scheme" default:"https" validate:"required,oneof=http https"`
	Host           string `json:"host" validate:"required"`
	Port           string `json:"port" validate:"required"`
	Path           string `json:"path" validate:"required"`
	Subdomain      string `json:"subdomain" validate:"required"`
	ExpectedOutput string `json:"expected_output" validate:"required"`
	Username       string `json:"username" validate:"required"`
	Password       string `json:"password" validate:"required"`
}

// ScoreCalDav checks if a caldav can be authenticated with username and password set and find the user
func ScoreCalDav(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	caldavproperties := &CalDavProperties{}
	err := json.Unmarshal(properties, &caldavproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %w", err))
		return
	}

	err = validate.Struct(caldavproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	authClient := webdav.HTTPClientWithBasicAuth(&http.Client{}, caldavproperties.Username, caldavproperties.Password)
	baseURL := ConstructURI(caldavproperties.Port, caldavproperties.Subdomain, caldavproperties.Host, caldavproperties.Path, caldavproperties.Scheme)
	client, err := webdav.NewClient(authClient, baseURL.String())
	if err != nil {
		ow.SetError(fmt.Errorf("unable to create client :%w", err))
		return
	}
	usr, err := client.FindCurrentUserPrincipal()
	if err != nil {
		ow.SetError(fmt.Errorf("unable to retrieve current user principal: %w", err))
		return
	}
	if caldavproperties.ExpectedOutput != "" && !strings.Contains(usr, caldavproperties.ExpectedOutput) {
		ow.SetError(fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, usr))
		return
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
