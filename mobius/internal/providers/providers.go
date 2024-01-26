package providers

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/config"
	"github.com/missingstudio/studio/backend/internal/providers/base"
	"github.com/missingstudio/studio/backend/internal/providers/openai"
)

type ProviderFactory interface {
	Create(token string) base.ProviderInterface
}

var providerFactories = make(map[string]ProviderFactory)

func init() {
	providerFactories["openai"] = openai.OpenAIProviderFactory{}
}

func GetProvider(ctx context.Context) (base.ProviderInterface, error) {
	providerName, ok := ctx.Value(config.ProviderKey{}).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("failed to get provider"))
	}

	authkey, ok := ctx.Value(config.AuthorizationKey{}).(string)
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("failed to get access key"))
	}

	providerFactory, ok := providerFactories[providerName]
	if !ok {
		return nil, connect.NewError(connect.CodeNotFound, errors.New("provider not found"))
	}
	return providerFactory.Create(authkey), nil
}
