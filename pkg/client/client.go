package client

import "ScoreTrak/pkg/team"

type ScoretrakClient struct {
	ip string
	port string
	token string
}

func NewScoretrakClient (ip, port, token string) ScoretrakClient{
	return ScoretrakClient{ip: ip, port: port, token: token}
}

func (s ScoretrakClient) GetAllTeams() []* team.Team{
	return nil
}