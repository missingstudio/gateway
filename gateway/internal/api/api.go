package api

import (
	"log/slog"
	"time"

	"github.com/missingstudio/ai/gateway/core/apikey"
	"github.com/missingstudio/ai/gateway/core/connection"
	"github.com/missingstudio/ai/gateway/core/prompt"
	"github.com/missingstudio/ai/gateway/core/provider"
	"github.com/missingstudio/ai/gateway/internal/ingester"
	iprovider "github.com/missingstudio/ai/gateway/internal/provider"
	"github.com/missingstudio/ai/gateway/internal/ratelimiter"
	"github.com/missingstudio/ai/gateway/pkg/database"
	"github.com/missingstudio/common/rest"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type API struct {
	Logger            *slog.Logger
	RestConfig        *rest.Config
	DBClient          *database.Client
	RestServer        *rest.Server
	RequestTimeout    time.Duration
	Ingester          ingester.Ingester
	RateLimiter       *ratelimiter.RateLimiter
	APIKeyService     *apikey.Service
	PromptService     *prompt.Service
	ProviderService   *provider.Service
	IProviderService  *iprovider.Service
	ConnectionService *connection.Service
}

func (api *API) Start() error {
	mux := api.routes()

	defer func() {
		api.Logger.Debug("cleaning up db")
		if err := api.DBClient.Close(); err != nil {
			api.Logger.Warn("db cleanup failed", "err", err)
		}
	}()

	server, err := rest.NewServer(api.RestConfig, api.Logger, mux)
	if err != nil {
		return err
	}

	api.RestServer = server
	// need this to make gRPC protocol work
	api.RestServer.Server.Handler = h2c.NewHandler(mux, &http2.Server{})
	return api.RestServer.Start()
}
