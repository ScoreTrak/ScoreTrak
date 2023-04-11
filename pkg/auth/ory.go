package auth

import (
	"context"

	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	client "github.com/ory/client-go"
)

func NewOryClient(ctx context.Context, c *config.Config) *client.APIClient {
	configuration := client.NewConfiguration()
	configuration.Servers = []client.ServerConfiguration{
		{
			URL: c.Auth.Ory.AdminApiUrl,
		},
	}

	apiClient := client.NewAPIClient(configuration)
	return apiClient
}
