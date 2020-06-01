package client

import (
	"net/http"
)

type ScoretrakClient struct {
	address    string
	port       string
	token      string
	httpClient *http.Client
}

func NewScoretrakClient(address, port, token string, client *http.Client) ScoretrakClient {
	return ScoretrakClient{address: address, port: port, token: token, httpClient: client}
}
