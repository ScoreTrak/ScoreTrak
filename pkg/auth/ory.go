package auth

import (
	"github.com/ScoreTrak/ScoreTrak/pkg/config"
	oclient "github.com/ory/client-go"
	kclient "github.com/ory/kratos-client-go"
)

func NewOryClient(c *config.Config) *oclient.APIClient {
	configuration := oclient.NewConfiguration()
	configuration.Servers = []oclient.ServerConfiguration{
		{
			URL: c.Auth.Ory.AdminApiUrl,
		},
	}
	kcfg := kclient.NewConfiguration()
	kcfg.Servers = []kclient.ServerConfiguration{
		{
			URL: c.Auth.Ory.AdminApiUrl,
		},
	}
	kApiClient := kclient.NewAPIClient(kcfg)
	kApiClient.GetConfig()

	apiClient := oclient.NewAPIClient(configuration)
	return apiClient
}
