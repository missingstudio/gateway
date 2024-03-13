package anyscale

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionProvider = &anyscaleProvider{}

type anyscaleProvider struct {
	info     base.Info
	config   base.Config
	provider provider.Provider
}

func (anyscale anyscaleProvider) Info() base.Info {
	return anyscale.info
}

func (anyscale anyscaleProvider) Config() base.Config {
	return anyscale.config
}

func (anyscale anyscaleProvider) Schema() []byte {
	return schema
}

func getAnyscaleInfo() base.Info {
	return base.Info{
		Title: "Anyscale",
		Name:  "anyscale",
		Description: `Anyscale Endpoints is a fast and scalable API to integrate OSS LLMs into your app.  
		Use our growing list of high performance models or deploy your own.`,
	}
}

func getAnyscaleConfig(baseURL string) base.Config {
	return base.Config{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["anyscale"] = func(provider provider.Provider) base.Provider {
		config := getAnyscaleConfig("https://api.endpoints.anyscale.com")
		return &anyscaleProvider{
			info:     getAnyscaleInfo(),
			config:   config,
			provider: provider,
		}
	}
}
