package azure

import (
	"context"
	"errors"
	"fmt"

	"github.com/missingstudio/studio/backend/internal/providers/openai"
	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (az *azureProvider) ChatCompletion(ctx context.Context, cr *llmv1.ChatCompletionRequest) (*llmv1.ChatCompletionResponse, error) {
	url := fmt.Sprintf(
		"https://%s.openai.azure.com/openai/deployments/%s%s?api-version=%s",
		az.ResourceName, az.DeploymentID, az.Config.ChatCompletions, az.APIVersion,
	)

	fmt.Println(url)
	return nil, errors.New("Not yet implemented")
}

func (az *azureProvider) GetModels() []string {
	return openai.OpenAIModels
}
