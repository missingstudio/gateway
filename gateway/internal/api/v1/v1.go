package v1

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/otelconnect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/missingstudio/studio/backend/internal/interceptor"
	"github.com/missingstudio/studio/common/middlewares"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"
)

type V1Handler struct {
	llmv1connect.UnimplementedLLMServiceHandler
}

func Register() (http.Handler, error) {
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return nil, fmt.Errorf("failed to create validate interceptor: %w", err)
	}

	otelconnectInterceptor, err := otelconnect.NewInterceptor(otelconnect.WithTrustRemote())
	if err != nil {
		return nil, fmt.Errorf("failed to create validate otel connect: %w", err)
	}

	compress1KB := connect.WithCompressMinBytes(1024)
	stdInterceptors := []connect.Interceptor{
		validateInterceptor,
		otelconnectInterceptor,
		interceptor.ProviderInterceptor(),
		interceptor.RetryInterceptor(),
	}

	services := []*vanguard.Service{
		vanguard.NewService(llmv1connect.NewLLMServiceHandler(
			&V1Handler{},
			compress1KB,
			connect.WithInterceptors(stdInterceptors...),
		)),
	}
	transcoderOptions := []vanguard.TranscoderOption{
		vanguard.WithUnknownHandler(middlewares.Custom404handler()),
	}

	return vanguard.NewTranscoder(services, transcoderOptions...)
}
