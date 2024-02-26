package executors

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/creasty/defaults"
	"github.com/digineo/go-ping"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome"
	"github.com/scoretrak/scoretrak/pkg/scorer/outcome_writer"
	"net"
	"strconv"
	"time"
)

type PingProperties struct {
	Host     string `json:"host" validate:"required"`
	Protocol string `json:"protocol" default:"ipv4" validate:"required"`
	Attempts string `json:"attempts" default:"3" valdate:"required"`
}

var ipv4opt = []string{"ipv4", "4", "ip4"}
var ipv6opt = []string{"ipv6", "6", "ip6"}

var ErrUnsupportedParameter = errors.New("invalid protocol selected")

//func (p *Ping) Validate() error {
//	_, err := strconv.Atoi(pingproperties.Attempts)
//	if err != nil {
//		return fmt.Errorf("unable to convert field attempts(%s) to int: %w", pingproperties.Attempts, err)
//	}
//	if ContainsString(ipv4opt, pingproperties.Protocol) || ContainsString(ipv6opt, pingproperties.Protocol) {
//		return nil
//	}
//	return fmt.Errorf("%w. Must be:'%s', '%s', '%s' for ipv4, or '%s', '%s', '%s' for ipv6", ErrUnsupportedParameter, ipv4opt[0], ipv4opt[1], ipv4opt[2], ipv6opt[0], ipv6opt[1], ipv4opt[2])
//}

func ScorePing(ctx context.Context, ow *outcome_writer.OutcomeWriter, properties []byte) {
	pingproperties := &PingProperties{}
	err := json.Unmarshal(properties, &pingproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("unable to unmarshall properties: %v", err))
		return
	}

	err = defaults.Set(pingproperties)
	if err != nil {
		ow.SetError(fmt.Errorf("default set error: %v", err))
		return
	}

	err = validate.Struct(pingproperties)
	if err != nil {
		// validationErrors := err.(validator.ValidationErrors)
		ow.SetError(fmt.Errorf("validation error: %w", err))
		return
	}

	var remoteAddr *net.IPAddr
	var pinger *ping.Pinger
	if ContainsString(ipv4opt, pingproperties.Protocol) {
		r, err := net.ResolveIPAddr("ip4", pingproperties.Host)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to resolve remote address: %w", err))
			return
		}
		remoteAddr = r
		p, err := ping.New("0.0.0.0", "")
		if err != nil {
			ow.SetError(fmt.Errorf("unable to initialize pinger, this is most likely a bug: %w", err))
			return
		}
		pinger = p
	} else {
		r, err := net.ResolveIPAddr("ip6", pingproperties.Host)
		if err != nil {
			ow.SetError(fmt.Errorf("unable to resolve remote address: %w", err))
			return
		}
		remoteAddr = r

		p, err := ping.New("", "::")
		if err != nil {
			ow.SetError(fmt.Errorf("unable to initialize pinger, this is most likely a bug: %w", err))
			return
		}
		pinger = p
	}
	defer pinger.Close()
	i, _ := strconv.Atoi(pingproperties.Attempts)
	_, err = pinger.PingAttempts(remoteAddr, time.Second*10, i) // TODO(thisisibrahimd) log round trip time
	if err != nil {
		ow.SetError(fmt.Errorf("unable to perform the ping: %w", err))
		return
	}

	ow.SetStatus(outcome.OUTCOME_STATUS_PASSED)
}
