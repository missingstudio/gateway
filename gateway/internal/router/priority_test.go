package router_test

import (
	"testing"

	"github.com/missingstudio/studio/backend/internal/mock"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/router"
	"github.com/stretchr/testify/require"
)

func TestPriorityRouter(t *testing.T) {
	type Provider struct {
		info base.ProviderInfo
	}

	type TestCase struct {
		providers           []Provider
		expectedProviderIDs []string
	}

	tests := map[string]TestCase{
		"openai": {[]Provider{
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

			routing := router.NewPriorityRouter(providers)
			iterator := routing.Iterator()

			for _, providerID := range tc.expectedProviderIDs {
				provider, err := iterator.Next()
				config := provider.Info()
				require.NoError(t, err)
				require.Equal(t, providerID, config.Name)
			}
		})
	}
}
