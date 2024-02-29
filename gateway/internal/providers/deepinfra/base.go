package deepinfra

import (
	_ "embed"

	"github.com/missingstudio/studio/backend/core/connection"
	"github.com/missingstudio/studio/backend/internal/providers/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionInterface = &deepinfraProvider{}

type deepinfraProvider struct {
	info   base.ProviderInfo
	config base.ProviderConfig
	conn   connection.Connection
}

func (anyscale deepinfraProvider) Info() base.ProviderInfo {
	return anyscale.info
}

func (deepinfra deepinfraProvider) Config() base.ProviderConfig {
	return deepinfra.config
}

func (deepinfra deepinfraProvider) Schema() []byte {
	return schema
}

func getDeepinfraInfo() base.ProviderInfo {
	return base.ProviderInfo{
		Title: "Deepinfra",
		Name:  "deepinfra",
		Description: `Deep Infra offers 100+ machine learning models from Text-to-Image, Object-Detection, 
		Automatic-Speech-Recognition, Text-to-Text Generation, and more!`,
	}
}

func getDeepinfraConfig(baseURL string) base.ProviderConfig {
	return base.ProviderConfig{
		BaseURL:         baseURL,
		ChatCompletions: "/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["deepinfra"] = func(connection connection.Connection) base.IProvider {
		config := getDeepinfraConfig("https://api.deepinfra.com/v1/openai")
		return &deepinfraProvider{
			info:   getDeepinfraInfo(),
			config: config,
			conn:   connection,
		}
	}
}
