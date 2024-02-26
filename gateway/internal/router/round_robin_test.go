package router_test

import (
	"testing"

	"github.com/missingstudio/studio/backend/internal/router"
	"github.com/stretchr/testify/require"
)

func TestRoundRobinRouter(t *testing.T) {
	type TestCase struct {
		providers        []router.RouterConfig
		expectedModelIDs []string
	}

	tests := map[string]TestCase{
		"public llms": {[]router.RouterConfig{
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

			iterator := router.NewRoundRobinRouter(providers)

			// loop three times over the whole pool to check if we return back to the begging of the list
			for _, providerName := range tc.expectedModelIDs {
				provider := iterator.Next()
				require.NotNil(t, provider)
				require.Equal(t, providerName, provider.Name)
			}
		})
	}
}
