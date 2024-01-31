package openai

import (
	"net/http"
	"strings"

	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/common/errors"
)

type OpenAIProviderFactory struct{}

func (f OpenAIProviderFactory) Create(headers http.Header) (base.ProviderInterface, error) {
	authorization := headers.Get(config.Authorization)
	if authorization == "" {
		return nil, errors.NewBadRequest("authorization header is required")
	}

	authorizationKey := strings.Replace(authorization, "Bearer ", "", 1)
	openAIProvider := NewOpenAIProvider(authorizationKey, "https://api.openai.com")
	return openAIProvider, nil
}

type OpenAIHeaders struct {
	APIKey string
}

type OpenAIProvider struct {
	Name   string
	Config base.ProviderConfig
	OpenAIHeaders
}

func NewOpenAIProvider(apikey string, baseURL string) *OpenAIProvider {
	config := getOpenAIConfig(baseURL)

	return &OpenAIProvider{
		Name: "Open AI",
		OpenAIHeaders: OpenAIHeaders{
			APIKey: apikey,
		},
		Config: config,
	}
}

func (oai *OpenAIProvider) GetName() string {
	return oai.Name
}

func getOpenAIConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}
