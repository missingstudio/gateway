package anyscale

import (
	"net/http"
	"strings"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/common/errors"
)

type AnyscaleProviderFactory struct{}

type AnyscaleHeaders struct {
	APIKey string `validate:"required" json:"Authorization" error:"API key is required"`
}

func (anyscale AnyscaleProviderFactory) GetHeaders(headers http.Header) (*AnyscaleHeaders, error) {
	var anyscaleHeaders AnyscaleHeaders
	if err := utils.UnmarshalHeader(headers, &anyscaleHeaders); err != nil {
		return nil, errors.New(err)
	}

	return &anyscaleHeaders, nil
}

func (anyscale AnyscaleProviderFactory) Create(headers http.Header) (base.ProviderInterface, error) {
	anyscaleHeaders, err := anyscale.GetHeaders(headers)
	if err != nil {
		return nil, err
	}

	anyscaleHeaders.APIKey = strings.Replace(anyscaleHeaders.APIKey, "Bearer ", "", 1)
	openAIProvider := NewAnyscaleProvider(*anyscaleHeaders)
	return openAIProvider, nil
}
