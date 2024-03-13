package provider

import (
	"github.com/missingstudio/ai/gateway/internal/provider/base"
	"github.com/missingstudio/common/errors"
	"github.com/xeipuuv/gojsonschema"
)

func Validate(provider base.Provider, data map[string]any) error {
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
