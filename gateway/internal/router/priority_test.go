package router_test

import (
	"testing"

	"github.com/missingstudio/studio/backend/internal/router"
	"github.com/stretchr/testify/require"
)

func TestPriorityRouter(t *testing.T) {
	type TestCase struct {
		providers           []router.RouterConfig
		expectedProviderIDs []string
	}

	tests := map[string]TestCase{
		"openai": {[]router.RouterConfig{
			{Name: "openai"},
			{Name: "anyscale"},
			{Name: "azure"},
		}, []string{"openai", "anyscale", "azure"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			providers := make([]router.RouterConfig, 0, len(tc.providers))

			for _, provider := range tc.providers {
				providers = append(providers, router.RouterConfig{Name: provider.Name})
			}

			iterator := router.NewPriorityRouter(providers)

			for _, providerID := range tc.expectedProviderIDs {
				provider, err := iterator.Next()
				require.NoError(t, err)
				require.Equal(t, providerID, provider.Name)
			}
		})
	}
}
