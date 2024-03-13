package deepinfra

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionProvider = &deepinfraProvider{}

type deepinfraProvider struct {
	info     base.Info
	config   base.Config
	provider provider.Provider
}

func (anyscale deepinfraProvider) Info() base.Info {
	return anyscale.info
}

func (deepinfra deepinfraProvider) Config() base.Config {
	return deepinfra.config
}

func (deepinfra deepinfraProvider) Schema() []byte {
	return schema
}

func getDeepinfraInfo() base.Info {
	return base.Info{
		Title: "Deepinfra",
		Name:  "deepinfra",
		Description: `Deep Infra offers 100+ machine learning models from Text-to-Image, Object-Detection, 
		Automatic-Speech-Recognition, Text-to-Text Generation, and more!`,
	}
}

func getDeepinfraConfig(baseURL string) base.Config {
	return base.Config{
		BaseURL:         baseURL,
		ChatCompletions: "/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["deepinfra"] = func(provider provider.Provider) base.Provider {
		config := getDeepinfraConfig("https://api.deepinfra.com/v1/openai")
		return &deepinfraProvider{
			info:     getDeepinfraInfo(),
			config:   config,
			provider: provider,
		}
	}
}
