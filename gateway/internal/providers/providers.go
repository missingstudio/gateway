package providers

import (
	"net/http"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/common/errors"
	"github.com/xeipuuv/gojsonschema"
)

type ProviderFactory interface {
	Create(headers http.Header) (base.IProvider, error)
}

func Validate(provider base.IProvider, data map[string]any) error {
	providerSchema := gojsonschema.NewBytesLoader(provider.Schema())
	connectionSchema := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(providerSchema, connectionSchema)
	if err != nil {
		return err
	}

	if !result.Valid() {
		var err error
		for _, desc := range result.Errors() {
			err = errors.NewBadRequest(desc.String())
		}

		if err != nil {
			return err
		}
	}

	return nil
}
