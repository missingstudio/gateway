package deepinfra

import (
	"net/http"
	"strings"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
)

type DeepinfraProviderFactory struct{}

type DeepinfraHeaders struct {
	APIKey string `validate:"required" json:"Authorization" error:"API key is required"`
}

func (deepinfra DeepinfraProviderFactory) GetHeaders(headers http.Header) (*DeepinfraHeaders, error) {
	var deepinfraHeaders DeepinfraHeaders
	if err := utils.UnmarshalHeader(headers, &deepinfraHeaders); err != nil {
		return nil, errors.New(err)
	}

	return &deepinfraHeaders, nil
}

func (deepinfra DeepinfraProviderFactory) Create(headers http.Header) (base.ProviderInterface, error) {
	deepinfraHeaders, err := deepinfra.GetHeaders(headers)
	if err != nil {
		return nil, err
	}

	deepinfraHeaders.APIKey = strings.Replace(deepinfraHeaders.APIKey, "Bearer ", "", 1)
	openAIProvider := NewDeepinfraProvider(*deepinfraHeaders)
	return openAIProvider, nil
}
