package openai

import (
	"net/http"
	"strings"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
)

type OpenAIProviderFactory struct{}

type OpenAIHeaders struct {
	APIKey string `validate:"required" json:"Authorization" error:"API key is required"`
}

func (oaif OpenAIProviderFactory) GetHeaders(headers http.Header) (*OpenAIHeaders, error) {
	var oaiHeaders OpenAIHeaders
	if err := utils.UnmarshalHeader(headers, &oaiHeaders); err != nil {
		return nil, errors.New(err)
	}

	return &oaiHeaders, nil
}

func (oaif OpenAIProviderFactory) Create(headers http.Header) (base.IProvider, error) {
	oaiHeaders, err := oaif.GetHeaders(headers)
	if err != nil {
		return nil, err
	}

	oaiHeaders.APIKey = strings.Replace(oaiHeaders.APIKey, "Bearer ", "", 1)
	openAIProvider := NewOpenAIProvider(*oaiHeaders)
	return openAIProvider, nil
}
