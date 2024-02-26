package auth

import (
	"github.com/scoretrak/scoretrak/pkg/config"
	client "github.com/ory/kratos-client-go"
)

func NewOryClient(c *config.Config) *client.APIClient {
	cfg := client.NewConfiguration()
	cfg.Servers = []client.ServerConfiguration{
		{
			URL: c.Auth.Ory.AdminApiUrl,
		},
	}
	apiClient := client.NewAPIClient(cfg)
	apiClient.GetConfig()

	return apiClient
}
