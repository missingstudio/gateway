package anyscale

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/models"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionInterface = &anyscaleProvider{}

type anyscaleProvider struct {
	info   base.ProviderInfo
	config base.ProviderConfig
	conn   models.Connection
}

func (anyscale anyscaleProvider) Info() base.ProviderInfo {
	return anyscale.info
}

func (anyscale anyscaleProvider) Config() base.ProviderConfig {
	return anyscale.config
}

func (anyscale anyscaleProvider) Schema() []byte {
	return schema
}

func getAnyscaleInfo() base.ProviderInfo {
	return base.ProviderInfo{
		Title: "Anyscale",
		Name:  "anyscale",
		Description: `Anyscale Endpoints is a fast and scalable API to integrate OSS LLMs into your app.  
		Use our growing list of high performance models or deploy your own.`,
	}
}

func getAnyscaleConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["anyscale"] = func(connection models.Connection) base.IProvider {
		config := getAnyscaleConfig("https://api.endpoints.anyscale.com")
		return &anyscaleProvider{
			info:   getAnyscaleInfo(),
			config: config,
			conn:   connection,
		}
	}
}
