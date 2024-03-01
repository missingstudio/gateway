package interceptor

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/core/apikey"
	"github.com/missingstudio/studio/backend/internal/constants"
	"github.com/missingstudio/studio/backend/internal/errors"
)

// NewAPIKeyInterceptor returns interceptor which is checking if api key exits
func NewAPIKeyInterceptor(logger *slog.Logger, aks *apikey.Service, authEnabled bool) connect.UnaryInterceptorFunc {
	return connect.UnaryInterceptorFunc(func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if !authEnabled {
				return next(ctx, req)
			}

			if authenticationSkipList[req.Spec().Procedure] {
				return next(ctx, req)
			}

			apiHeader := req.Header().Get(constants.XMSAPIKey)
			if apiHeader == "" {
				logger.Info("request without api key",
					"api_key", apiHeader,
					"addr", req.Peer().Addr,
					"endpoint", req.Spec().Procedure)
				return nil, errors.ErrUnauthenticated
			}

			apikey := req.Header().Get(constants.XMSAPIKey)
			if _, err := aks.GetByToken(context.Background(), apikey); err != nil {
				return nil, errors.ErrUnauthenticated
			}

			return next(ctx, req)
		})
	})
}

// authenticationSkipList stores path to skip authentication, by default its enabled for all requests
var authenticationSkipList = map[string]bool{
	"/llm.v1.LLMService/ListModels":        true,
	"/llm.v1.LLMService/ListProviders":     true,
	"/llm.v1.LLMService/GetProviderConfig": true,
	"/llm.v1.LLMService/ListAPIKeys":       true,
	"/llm.v1.LLMService/CreateAPIKey":      true,
}
