package azure

import (
	"context"
	"errors"
	"fmt"

	llmv1 "github.com/missingstudio/studio/protos/pkg/llm"
)

func (az *AzureProvider) ChatCompilation(ctx context.Context, cr *llmv1.CompletionRequest) (*llmv1.CompletionResponse, error) {
	url := fmt.Sprintf(
		"https://%s.openai.azure.com/openai/deployments/%s%s?api-version=%s",
		az.ResourceName, az.DeploymentID, az.Config.ChatCompletions, az.APIVersion,
	)

	fmt.Println(url)
	return nil, errors.New("Not yet implemented")
}
