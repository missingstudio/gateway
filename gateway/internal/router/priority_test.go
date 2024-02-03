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
		Name string
	}

	type TestCase struct {
		providers        []Provider
		expectedModelIDs []string
	}

	tests := map[string]TestCase{
		"openai": {[]Provider{{"openai"}, {"anyscale"}, {"azure"}}, []string{"openai", "anyscale", "azure"}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			providers := make([]base.ProviderInterface, 0, len(tc.providers))

			for _, provider := range tc.providers {
				providers = append(providers, mock.NewProviderMock(provider.Name))
			}

			routing := router.NewPriorityRouter(providers)
			iterator := routing.Iterator()

			for _, modelID := range tc.expectedModelIDs {
				model, err := iterator.Next()
				require.NoError(t, err)
				require.Equal(t, modelID, model.GetName())
			}
		})
	}
}
