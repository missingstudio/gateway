package azure

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/internal/provider/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionProvider = &azureProvider{}

type azureProvider struct {
	info   base.Info
	config base.Config
}

func (anyscale azureProvider) Info() base.Info {
	return anyscale.info
}

func (az azureProvider) Config() base.Config {
	return az.config
}

func (az azureProvider) Schema() []byte {
	return schema
}

func GetAzureInfo() base.Info {
	return base.Info{
		Title:       "Azure",
		Name:        "azure",
		Description: "Azure OpenAI Service offers industry-leading coding and language AI models that you can fine-tune to your specific needs for a variety of use cases.",
	}
}

func GetAzureConfig() base.Config {
	return base.Config{
		BaseURL:         "",
		ChatCompletions: "/chat/completions",
	}
}

func init() {}
