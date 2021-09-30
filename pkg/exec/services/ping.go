package services

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/digineo/go-ping"
)

type Ping struct {
	Protocol string
	Attempts string
}

var ipv4opt = []string{"ipv4", "4", "ip4"}
var ipv6opt = []string{"ipv6", "6", "ip6"}

func NewPing() *Ping {
	f := Ping{Protocol: "ipv4", Attempts: "1"}
	return &f
}

var ErrUnsupportedParameter = errors.New("invalid protocol selected")

func (p *Ping) Validate() error {
	_, err := strconv.Atoi(p.Attempts)
	if err != nil {
		return fmt.Errorf("unable to convert field attempts(%s) to int: %w", p.Attempts, err)
	}
	if ContainsString(ipv4opt, p.Protocol) || ContainsString(ipv6opt, p.Protocol) {
		return nil
	}
	return fmt.Errorf("%w. Must be:'%s', '%s', '%s' for ipv4, or '%s', '%s', '%s' for ipv6", ErrUnsupportedParameter, ipv4opt[0], ipv4opt[1], ipv4opt[2], ipv6opt[0], ipv6opt[1], ipv4opt[2])
}

func (p *Ping) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	var remoteAddr *net.IPAddr
	var pinger *ping.Pinger
	if ContainsString(ipv4opt, p.Protocol) {
		r, err := net.ResolveIPAddr("ip4", e.Host)
		if err != nil {
			return false, "", fmt.Errorf("unable to resolve remote address: %w", err)
		}
		remoteAddr = r
		p, err := ping.New("0.0.0.0", "")
		if err != nil {
			return false, "", fmt.Errorf("unable to initialize pinger, this is most likely a bug: %w", err)
		}
		pinger = p
	} else {
		r, err := net.ResolveIPAddr("ip6", e.Host)
		if err != nil {
			return false, "", fmt.Errorf("unable to resolve remote address: %w", err)
		}
		remoteAddr = r

		p, err := ping.New("", "::")
		if err != nil {
			return false, "", fmt.Errorf("unable to initialize pinger, this is most likely a bug: %w", err)
		}
		pinger = p
	}
	defer pinger.Close()
	i, _ := strconv.Atoi(p.Attempts)
	rtt, err := pinger.PingAttempts(remoteAddr, time.Until(e.Deadline())/time.Duration(i), i)
	if err != nil {
		return false, "", fmt.Errorf("unable to perform the ping: %w", err)
	}
	return true, fmt.Sprintf("%s\nRound trip time: %s", Success, rtt.String()), nil
}

// Below Code has some very nasty errors that are in the underlying library(For instance: https://github.com/sparrc/go-ping/pull/80). Until they are fixed, we will use https://github.com/digineo/go-ping
// func (p *Ping) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
//	pinger, err := ping.NewPinger(e.Host)
//	if err != nil {
//		return false, "Unable to initialize new pinger", err
//	}
//	pinger.Timeout = time.Until(e.Deadline())
//	pinger.SetPrivileged(true)
//	i, err := strconv.Atoi(p.Count)
//	if err != nil {
//		return false, "Unable to convert Count to int", err
//	}
//	pinger.Count = i
//	pinger.Run()
//	stats := pinger.Statistics()
//	if stats.PacketLoss != 0 {
//		return false, fmt.Sprintf("Packet loss was not 0%%, instead it was: %.2f%%", stats.PacketLoss), nil
//	}
//	return true, Success, nil
// }
