package v1

import (
	"context"

	"connectrpc.com/connect"
	"github.com/missingstudio/ai/common/errors"
	llmv1 "github.com/missingstudio/ai/protos/pkg/llm/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *V1Handler) ListTrackingLogs(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[llmv1.LogResponse], error) {
	response, err := s.ingester.Get()
	if err != nil {
		return nil, errors.New(err)
	}

	logs := []*structpb.Struct{}
	for _, log := range response {
		point := map[string]any{
			"latency":           log["latency"],
			"model":             log["model"],
			"provider":          log["provider"],
			"total_tokens":      log["total_tokens"],
			"prompt_tokens":     log["prompt_tokens"],
			"completion_tokens": log["completion_tokens"],
		}
		pointdata, _ := structpb.NewStruct(point)
		logs = append(logs, pointdata)
	}

	return connect.NewResponse(&llmv1.LogResponse{
		Logs: logs,
	}), nil
}
