package v1

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/missingstudio/studio/backend/internal/api"
	"github.com/missingstudio/studio/backend/internal/connections"
	"github.com/missingstudio/studio/backend/internal/ingester"
	"github.com/missingstudio/studio/backend/internal/interceptor"
	"github.com/missingstudio/studio/backend/internal/providers"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"
)

type V1Handler struct {
	llmv1connect.UnimplementedLLMServiceHandler
	ingester          ingester.Ingester
	providerService   *providers.Service
	connectionService *connections.Service
}

func NewHandlerV1(d *api.Deps) *V1Handler {
	return &V1Handler{
		ingester:          d.Ingester,
		providerService:   d.ProviderService,
		connectionService: d.ConnectionService,
	}
}

func Register(d *api.Deps) (http.Handler, error) {
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return nil, fmt.Errorf("failed to create validate interceptor: %w", err)
	}

	v1Handler := NewHandlerV1(d)
	otelconnectInterceptor, err := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	if err != nil {
		return nil, fmt.Errorf("failed to create validate otel connect: %w", err)
	}

	compress1KB := connect.WithCompressMinBytes(1024)
	stdInterceptors := []connect.Interceptor{
		validateInterceptor,
		otelconnectInterceptor,
		interceptor.NewLoggingInterceptor(d.Logger),
		interceptor.RateLimiterInterceptor(d.RateLimiter),
		interceptor.RetryInterceptor(),
	}

	services := []*vanguard.Service{
		vanguard.NewService(llmv1connect.NewLLMServiceHandler(
			v1Handler,
			compress1KB,
			connect.WithInterceptors(stdInterceptors...),
		)),
	}
	transcoderOptions := []vanguard.TranscoderOption{
		vanguard.WithUnknownHandler(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			http.Error(w, "custom 404 error", http.StatusNotFound)
		})),
	}

	return vanguard.NewTranscoder(services, transcoderOptions...)
}
