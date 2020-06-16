package client

import (
	"ScoreTrak/pkg/report"
	"fmt"
)

type reportClient struct {
	s ScoretrakClient
}

func NewReportClient(c ScoretrakClient) report.Serv {
	return &reportClient{c}
}

func (c reportClient) Get() (*report.Report, error) {
	conf := &report.Report{}
	err := genericGet(conf, fmt.Sprintf("/report"), c.s)
	if err != nil {
		return nil, err
	}
	return conf, nil
}