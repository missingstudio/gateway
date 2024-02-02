package v1

import (
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/validate"
	"connectrpc.com/vanguard"
	"github.com/missingstudio/studio/backend/pkg/utils"
	"github.com/missingstudio/studio/protos/pkg/llm/llmv1connect"
)

type V1Handler struct {
	llmv1connect.UnimplementedLLMServiceHandler
}

func Register() (http.Handler, error) {
	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return nil, fmt.Errorf("validate interceptor not created: %w", err)
	}

	compress1KB := connect.WithCompressMinBytes(1024)
	services := []*vanguard.Service{
		vanguard.NewService(llmv1connect.NewLLMServiceHandler(
			&V1Handler{},
			compress1KB,
			connect.WithInterceptors(validateInterceptor, utils.ProviderInterceptor()),
		)),
	}
	transcoderOptions := []vanguard.TranscoderOption{
		vanguard.WithUnknownHandler(utils.Custom404handler()),
	}

	return vanguard.NewTranscoder(services, transcoderOptions...)
}
