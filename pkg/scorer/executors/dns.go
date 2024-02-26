package executors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bogdanovich/dns_resolver"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"golang.org/x/exp/slices"
)

type DNSProperties struct {
	RecordType     string `json:"record_type" validate:"required,oneof=A AAAA CNAME"`
	Host           string `json:"host" validate:"required,hostname"`
	Lookup         string `json:"lookup" validate:"required,hostname"`
	ExpectedOutput string `json:"expected_output" validate:"required"`
}

var ErrDNSLookupMissing = errors.New("you must pass a lookup parameter for DNS test")

func ScoreDns(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	dnsproperties := &DNSProperties{}
	err := json.Unmarshal(properties, &dnsproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %w", err))
		return
	}

	err = validate.Struct(dnsproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	resolver := dns_resolver.New([]string{dnsproperties.Host})
	resolver.RetryTimes = 1
	ips, err := resolver.LookupHost(dnsproperties.Lookup)
	if err != nil {
		ow.SetError(fmt.Errorf("encountered an error while looking up the host: %w", err))
		return
	}

	if dnsproperties.ExpectedOutput != "" {
		ipFound := slices.Contains(IPToStringSlice(ips), dnsproperties.ExpectedOutput)

		if !ipFound {
			ow.SetError(fmt.Errorf("Expected output is not received"))
			return
		}
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}

// Todo(thisisibrahimd): Substitute this DNS library with one that is more flexible (Support lookup types other than A)
