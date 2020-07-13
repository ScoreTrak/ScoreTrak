package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/report"
)

type reportClient struct {
	s ScoretrakClient
}

func NewReportClient(c ScoretrakClient) report.Serv {
	return &reportClient{c}
}

func (c reportClient) Get() (*report.Report, error) {
	conf := &report.Report{}
	err := c.s.genericGet(conf, fmt.Sprintf("/report"))
	if err != nil {
		return nil, err
	}
	return conf, nil
}
