package services

import (
	"ScoreTrak/pkg/exec"
	"fmt"
	"github.com/sparrc/go-ping"
	"strconv"
	"time"
)

type Ping struct {
	Count string `json:"count"`
}

func NewPing() *Ping {
	f := Ping{Count: "3"}
	return &f
}

func (p *Ping) Validate() error {
	return nil
}

func (p *Ping) Execute(e exec.Exec) (passed bool, log string, err error) {
	pinger, err := ping.NewPinger(e.Host)
	if err != nil {
		return false, "Unable to initialize new pinger", err
	}
	pinger.Timeout = time.Until(e.Deadline())
	pinger.SetPrivileged(true)
	i, err := strconv.Atoi(p.Count)
	if err != nil {
		return false, "Unable to convert Count to int", err
	}
	pinger.Count = i
	pinger.Run()
	stats := pinger.Statistics()
	if stats.PacketLoss != 0 {
		return false, fmt.Sprintf("Packet loss was not 0, instead it was: %d", stats.PacketLoss), nil
	}
	return true, "Success!", nil
}
