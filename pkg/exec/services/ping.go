package services

import (
	"fmt"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/digineo/go-ping"
	"net"
	"strconv"
	"time"
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

func (p *Ping) Validate() error {
	if exec.ContainsString(ipv4opt, p.Protocol) || exec.ContainsString(ipv6opt, p.Protocol) {
		return nil
	}
	return fmt.Errorf("protocol parameter should either be '%s' '%s' '%s' for ipv4, or '%s' '%s' '%s' for ipv6", ipv4opt[0], ipv4opt[1], ipv4opt[2], ipv6opt[0], ipv6opt[1], ipv4opt[2])
}

func (p *Ping) Execute(e exec.Exec) (passed bool, log string, err error) {
	var remoteAddr *net.IPAddr
	var pinger *ping.Pinger
	if exec.ContainsString(ipv4opt, p.Protocol) {
		if r, err := net.ResolveIPAddr("ip4", e.Host); err != nil {
			return false, "Unable to resolve remote address", err
		} else {
			remoteAddr = r
		}
		if p, err := ping.New("0.0.0.0", ""); err != nil {
			return false, "Unable to initialize pinger, this is most likely a bug", err
		} else {
			pinger = p
		}
	} else {
		if r, err := net.ResolveIPAddr("ip6", e.Host); err != nil {
			return false, "Unable to resolve remote address", err
		} else {
			remoteAddr = r
		}
		if p, err := ping.New("", "::"); err != nil {
			return false, "Unable to initialize pinger, this is most likely a bug", err
		} else {
			pinger = p
		}
	}
	defer pinger.Close()
	i, err := strconv.Atoi(p.Attempts)
	if err != nil {
		return false, fmt.Sprintf("Unable to convert %s to int", p.Attempts), err
	}
	rtt, err := pinger.PingAttempts(remoteAddr, time.Until(e.Deadline())/time.Duration(i), i)
	if err != nil {
		return false, "Unable to perform the ping", err
	}
	return true, fmt.Sprintf("Success!\nRound trip time: %s", rtt.String()), nil
}

//Below Code has some very nasty errors that are in the underlying library(For instance: https://github.com/sparrc/go-ping/pull/80). Until they are fixed, we will use https://github.com/digineo/go-ping
//func (p *Ping) Execute(e exec.Exec) (passed bool, log string, err error) {
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
//	return true, "Success!", nil
//}
