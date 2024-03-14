package v1

import (
	"github.com/missingstudio/ai/gateway/core/apikey"
	"github.com/missingstudio/ai/gateway/core/prompt"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/ingester"
	iprovider "github.com/missingstudio/ai/gateway/internal/provider"
	"github.com/missingstudio/protos/pkg/llm/v1/llmv1connect"
	"github.com/missingstudio/protos/pkg/prompt/v1/promptv1connect"
)

type V1Handler struct {
	llmv1connect.UnimplementedLLMServiceHandler
	promptv1connect.UnimplementedPromptRegistryServiceHandler
	ingester         ingester.Ingester
	apikeyService    *apikey.Service
	promptService    *prompt.Service
	providerService  *provider.Service
	iProviderService *iprovider.Service
}

func NewHandlerV1(
	ingester ingester.Ingester,
	apikeyService *apikey.Service,
	promptService *prompt.Service,
	providerService *provider.Service,
	iProviderService *iprovider.Service,
) *V1Handler {
	return &V1Handler{
		ingester:         ingester,
		apikeyService:    apikeyService,
		promptService:    promptService,
		providerService:  providerService,
		iProviderService: iProviderService,
	}
}
