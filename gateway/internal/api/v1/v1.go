package v1

import (
	"github.com/missingstudio/ai/gateway/core/apikey"
	"github.com/missingstudio/ai/gateway/core/connection"
	"github.com/missingstudio/ai/gateway/core/prompt"
	"github.com/missingstudio/ai/gateway/internal/ingester"
	"github.com/missingstudio/ai/gateway/internal/providers"
	"github.com/missingstudio/ai/protos/pkg/llm/v1/llmv1connect"
	"github.com/missingstudio/ai/protos/pkg/prompt/v1/promptv1connect"
)

type V1Handler struct {
	llmv1connect.UnimplementedLLMServiceHandler
	promptv1connect.UnimplementedPromptRegistryServiceHandler
	ingester          ingester.Ingester
	providerService   *providers.Service
	connectionService *connection.Service
	apikeyService     *apikey.Service
	promptService     *prompt.Service
}

func NewHandlerV1(
	ingester ingester.Ingester,
	providerService *providers.Service,
	connectionService *connection.Service,
	apikeyService *apikey.Service,
	promptService *prompt.Service,
) *V1Handler {
	return &V1Handler{
		ingester:          ingester,
		providerService:   providerService,
		connectionService: connectionService,
		apikeyService:     apikeyService,
		promptService:     promptService,
	}
}
