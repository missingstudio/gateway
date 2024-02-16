package togetherai

import (
	"net/http"
	"strings"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
)

type TogetherAIProviderFactory struct{}

type TogetherAIHeaders struct {
	APIKey string `validate:"required" json:"Authorization" error:"API key is required"`
}

func (ta TogetherAIProviderFactory) GetHeaders(headers http.Header) (*TogetherAIHeaders, error) {
	var togetherAIHeaders TogetherAIHeaders
	if err := utils.UnmarshalHeader(headers, &togetherAIHeaders); err != nil {
		return nil, errors.New(err)
	}

	return &togetherAIHeaders, nil
}

func (ta TogetherAIProviderFactory) Create(headers http.Header) (base.ProviderInterface, error) {
	togetherAIHeaders, err := ta.GetHeaders(headers)
	if err != nil {
		return nil, err
	}

	togetherAIHeaders.APIKey = strings.Replace(togetherAIHeaders.APIKey, "Bearer ", "", 1)
	openAIProvider := NewTogetherAIProvider(*togetherAIHeaders)
	return openAIProvider, nil
}
