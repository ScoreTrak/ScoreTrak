package services

import (
	"errors"
	"fmt"
	"github.com/bogdanovich/dns_resolver"

	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
)

type DNS struct {
	Lookup         string
	ExpectedOutput string
}

func NewDNS() *DNS {
	f := DNS{}
	return &f
}

var ErrDNSLookupMissing = errors.New("you must pass a lookup parameter for DNS test")

func (p *DNS) Validate() error {
	if p.Lookup != "" {
		return nil
	}
	return ErrDNSLookupMissing
}

func (p *DNS) Execute(e exec.Exec) (passed bool, logOutput string, err error) {
	resolver := dns_resolver.New([]string{e.HostAddress})
	resolver.RetryTimes = 1
	ip, err := resolver.LookupHost(p.Lookup)
	if err != nil {
		return false, "", fmt.Errorf("encountered an error while looking up the host: %w", err)
	}
	if p.ExpectedOutput != "" && ip[0].String() != p.ExpectedOutput {
		return false, "", fmt.Errorf("%w. Output Received: %s", ErrDidNotMatchExpectedOutput, ip[0].String())
	}
	return true, Success, nil
}

// Todo: Substitute this DNS library with one that is more flexible (Support lookup types other than A)
