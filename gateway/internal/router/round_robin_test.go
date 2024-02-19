package router_test

import (
	"testing"

	"github.com/missingstudio/studio/backend/internal/mock"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/router"
	"github.com/stretchr/testify/require"
)

func TestRoundRobinRouter(t *testing.T) {
	type Provider struct {
		info base.ProviderInfo
	}

	type TestCase struct {
		providers        []Provider
		expectedModelIDs []string
	}

	tests := map[string]TestCase{
		"public llms": {[]Provider{
			{info: base.ProviderInfo{Name: "openai"}},
			{info: base.ProviderInfo{Name: "anyscale"}},
			{info: base.ProviderInfo{Name: "azure"}},
		}, []string{"openai", "anyscale", "azure"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			providers := make([]base.IProvider, 0, len(tc.providers))

			for _, provider := range tc.providers {
				providers = append(providers, mock.NewProviderMock(provider.info.Name))
			}

			routing := router.NewRoundRobinRouter(providers)
			iterator := routing.Iterator()

			// loop three times over the whole pool to check if we return back to the begging of the list
			for _, providerName := range tc.expectedModelIDs {
				provider, err := iterator.Next()
				config := provider.Info()
				require.NoError(t, err)
				require.Equal(t, providerName, config.Name)
			}
		})
	}
}
