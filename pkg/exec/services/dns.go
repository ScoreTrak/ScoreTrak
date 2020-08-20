package services

import (
	"errors"
	"github.com/ScoreTrak/ScoreTrak/pkg/exec"
	"github.com/bogdanovich/dns_resolver"
)

type DNS struct {
	Lookup         string `json:"lookup"`
	ExpectedOutput string `json:"expected_output"`
}

func NewDNS() *DNS {
	f := DNS{}
	return &f
}

func (p *DNS) Validate() error {
	if p.Lookup != "" {
		return nil
	}
	return errors.New("you must pass a lookup parameter for DNS test")
}

func (p *DNS) Execute(e exec.Exec) (passed bool, log string, err error) {
	resolver := dns_resolver.New([]string{e.Host})
	resolver.RetryTimes = 1
	ip, err := resolver.LookupHost(p.Lookup)
	if err != nil {
		return false, "Encountered an error while looking up the host", err
	}
	if p.ExpectedOutput != "" && ip[0].String() != p.ExpectedOutput {
		return false, "Expected output did not match", nil //TODO: Make a more meaningful output
	}
	return true, "Success!", nil
}

//Todo: Substitute this DNS library with one that is more flexible (Support lookup types other than A)
