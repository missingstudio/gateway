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
		Name string
	}

	type TestCase struct {
		providers        []Provider
		expectedModelIDs []string
	}

	tests := map[string]TestCase{
		"public llms": {[]Provider{{"openai"}, {"anyscale"}, {"azure"}}, []string{"openai", "anyscale", "azure"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			providers := make([]base.ProviderInterface, 0, len(tc.providers))

			for _, provider := range tc.providers {
				providers = append(providers, mock.NewProviderMock(provider.Name))
			}

			routing := router.NewRoundRobinRouter(providers)
			iterator := routing.Iterator()

			// loop three times over the whole pool to check if we return back to the begging of the list
			for _, providerName := range tc.expectedModelIDs {
				provider, err := iterator.Next()
				require.NoError(t, err)
				require.Equal(t, providerName, provider.GetName())
			}
		})
	}
}
