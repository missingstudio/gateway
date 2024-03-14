package groq

import (
	_ "embed"

	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/provider/base"
)

//go:embed schema.json
var schema []byte

var _ base.ChatCompletionProvider = &groqProvider{}

type groqProvider struct {
	info     base.Info
	config   base.Config
	provider provider.Provider
}

func (groq groqProvider) Info() base.Info {
	return groq.info
}

func (groq groqProvider) Config() base.Config {
	return groq.config
}

func (groq groqProvider) Schema() []byte {
	return schema
}

func getGroqInfo() base.Info {
	return base.Info{
		Title: "Groq",
		Name:  "groq",
		Description: `Groq Endpoints is a fast and scalable API to integrate OSS LLMs into your app.  
		Use our growing list of high performance models or deploy your own.`,
	}
}

func getGroqConfig(baseURL string) base.Config {
	return base.Config{
		BaseURL:         baseURL,
		ChatCompletions: "/v1/chat/completions",
	}
}

func init() {
	base.ProviderRegistry["groq"] = func(provider provider.Provider) base.Provider {
		config := getGroqConfig("https://api.groq.com/openai")
		return &groqProvider{
			info:     getGroqInfo(),
			config:   config,
			provider: provider,
		}
	}
}
