package api

import (
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/go-chi/chi/v5"
	v1 "github.com/missingstudio/ai/gateway/internal/api/v1"
	"github.com/missingstudio/ai/gateway/internal/interceptor"
	"github.com/missingstudio/protos/pkg/llm/v1/llmv1connect"
	"github.com/missingstudio/protos/pkg/prompt/v1/promptv1connect"
)

func (api *API) routes() *chi.Mux {
	router := chi.NewRouter()
	v1Handler := v1.NewHandlerV1(
		api.Ingester,
		api.APIKeyService,
		api.PromptService,
		api.ProviderService,
		api.IProviderService,
	)
	compress1KB := connect.WithCompressMinBytes(1024)
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		log.Fatalf("failed to validate interceptor: %v", err)
	}

	otelconnectInterceptor, err := otelconnect.NewInterceptor(
		otelconnect.WithoutServerPeerAttributes(),
		otelconnect.WithTrustRemote(),
	)
	if err != nil {
		log.Fatalf("failed to create open telemetry interceptor: %v", err)
	}

	stdInterceptors := []connect.Interceptor{
		validateInterceptor,
		otelconnectInterceptor,
		interceptor.NewAPIKeyInterceptor(api.Logger, api.APIKeyService, false),
		interceptor.WithHeaderConfig(),
		interceptor.WithTimeout(api.RequestTimeout),
		interceptor.RateLimiterInterceptor(api.RateLimiter),
		interceptor.RetryInterceptor(),
		interceptor.NewLoggingInterceptor(api.Logger),
	}

	services := []*vanguard.Service{
		vanguard.NewService(llmv1connect.NewLLMServiceHandler(
			v1Handler,
			compress1KB,
			connect.WithInterceptors(stdInterceptors...),
		)),
		vanguard.NewService(promptv1connect.NewPromptRegistryServiceHandler(
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

	vanguardHandler, err := vanguard.NewTranscoder(services, transcoderOptions...)
	if err != nil {
		log.Fatalf("failed to create vanguard transcoder: %v", err)
	}

	router.Mount("/", vanguardHandler)
	router.Mount(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(
			llmv1connect.LLMServiceName,
			promptv1connect.PromptRegistryServiceName,
		),
		compress1KB,
	))

	reflector := grpcreflect.NewStaticReflector(
		llmv1connect.LLMServiceName,
		promptv1connect.PromptRegistryServiceName,
	)
	router.Mount(grpcreflect.NewHandlerV1(reflector, compress1KB))
	router.Mount(grpcreflect.NewHandlerV1Alpha(reflector, compress1KB))
	return router
}
