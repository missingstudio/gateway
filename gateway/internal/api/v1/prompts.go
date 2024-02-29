package v1

import (
	"bytes"
	"context"
	"text/template"

	"connectrpc.com/connect"
	"github.com/missingstudio/studio/backend/core/prompt"
	"github.com/missingstudio/studio/common/errors"
	promptv1 "github.com/missingstudio/studio/protos/pkg/prompt/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *V1Handler) ListPrompts(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[promptv1.ListPromptsResponse], error) {
	prompts, err := s.promptService.GetAll(ctx)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	data := []*promptv1.Prompt{}
	for _, p := range prompts {
		pmetadata, _ := structpb.NewStruct(p.Metadata)

		data = append(data, &promptv1.Prompt{
			Id:          p.ID.String(),
			Name:        p.Name,
			Description: p.Description,
			Template:    p.Template,
			Metadata:    pmetadata,
		})
	}

	return connect.NewResponse(&promptv1.ListPromptsResponse{
		Prompt: data,
	}), nil
}

func (s *V1Handler) CreatePrompt(ctx context.Context, req *connect.Request[promptv1.CreatePromptRequest]) (*connect.Response[promptv1.CreatePromptResponse], error) {
	prompt := prompt.Prompt{
		Name:        req.Msg.Name,
		Description: req.Msg.Description,
		Template:    req.Msg.Template,
		Metadata:    req.Msg.Metadata.AsMap(),
	}

	prompt, err := s.promptService.Upsert(ctx, prompt)
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	stMetadata, err := structpb.NewStruct(prompt.Metadata)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	return connect.NewResponse(&promptv1.CreatePromptResponse{
		Name:        prompt.Name,
		Description: prompt.Description,
		Template:    prompt.Template,
		Metadata:    stMetadata,
	}), nil
}

func (s *V1Handler) GetPrompt(ctx context.Context, req *connect.Request[promptv1.GetPromptRequest]) (*connect.Response[promptv1.GetPromptResponse], error) {
	prompt, err := s.promptService.GetByName(ctx, req.Msg.Name)
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	stMetadata, err := structpb.NewStruct(prompt.Metadata)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	p := &promptv1.Prompt{
		Id:          prompt.ID.String(),
		Name:        prompt.Name,
		Description: prompt.Description,
		Template:    prompt.Template,
		Metadata:    stMetadata,
	}

	return connect.NewResponse(&promptv1.GetPromptResponse{
		Prompt: p,
	}), nil
}

func (s *V1Handler) GetPromptValue(ctx context.Context, req *connect.Request[promptv1.GetPromptValueRequest]) (*connect.Response[promptv1.GetPromptValueResponse], error) {
	p, err := s.promptService.GetByName(ctx, req.Msg.Name)
	if err != nil {
		return nil, errors.NewNotFound(err.Error())
	}

	var buf bytes.Buffer
	tmpl := template.Must(template.New("prompt").Parse(p.Template))
	err = tmpl.Execute(&buf, req.Msg.Data.AsMap())
	if err != nil {
		return nil, errors.New(err)
	}

	return connect.NewResponse(&promptv1.GetPromptValueResponse{
		Data: buf.String(),
	}), nil
}
