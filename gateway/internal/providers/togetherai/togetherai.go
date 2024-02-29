package togetherai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/missingstudio/studio/backend/core/chat"
	"github.com/missingstudio/studio/backend/core/connection"
	"github.com/missingstudio/studio/backend/pkg/requester"
)

func (ta *togetherAIProvider) ChatCompletion(ctx context.Context, payload *chat.ChatCompletionRequest) (*chat.ChatCompletionResponse, error) {
	client := requester.NewHTTPClient()

	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal openai chat request payload: %w", err)
	}

	requestURL := fmt.Sprintf("%s%s", ta.config.BaseURL, ta.config.ChatCompletions)
	req, err := http.NewRequestWithContext(ctx, "POST", requestURL, bytes.NewReader(rawPayload))
	if err != nil {
		return nil, err
	}

	req = ta.AddDefaultHeaders(req, connection.AuthorizationHeader)
	resp, err := client.SendRequestRaw(req)
	if err != nil {
		return nil, err
	}

	data := &chat.ChatCompletionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		return nil, err
	}

	return data, nil
}

func (ta *togetherAIProvider) AddDefaultHeaders(req *http.Request, key string) *http.Request {
	connectionConfigMap := ta.conn.GetConfig([]string{key})

	var authorizationHeader string
	if val, ok := connectionConfigMap[connection.AuthorizationHeader].(string); ok && val != "" {
		authorizationHeader = val
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authorizationHeader)
	return req
}

func (*togetherAIProvider) Models() []string {
	return []string{
		"togethercomputer/CodeLlama-7b-Instruct",
		"togethercomputer/CodeLlama-13b-Instruct",
		"togethercomputer/CodeLlama-34b-Instruct",
		"codellama/CodeLlama-70b-Instruct-hf",
		"togethercomputer/llama-2-7b-chat",
		"togethercomputer/llama-2-13b-chat",
		"togethercomputer/llama-2-70b-chat",
		"mistralai/Mistral-7B-Instruct-v0.1",
		"mistralai/Mixtral-8x7B-Instruct-v0.1",
		"Phind/Phind-CodeLlama-34B-v2",
		"WizardLM/WizardCoder-Python-34B-V1.0",

		"zero-one-ai/Yi-34B-Chat",
		"Austism/chronos-hermes-13b",
		"deepseek-ai/deepseek-coder-33b-instruct",
		"garage-bAInd/Platypus2-70B-instruct",
		"Gryphe/MythoMax-L2-13b",
		"lmsys/vicuna-13b-v1.5",
		"lmsys/vicuna-7b-v1.5",
		"codellama/CodeLlama-13b-Instruct-hf",
		"codellama/CodeLlama-34b-Instruct-hf",
		"codellama/CodeLlama-70b-Instruct-hf",
		"codellama/CodeLlama-7b-Instruct-hf",
		"meta-llama/Llama-2-70b-chat-hf",
		"meta-llama/Llama-2-13b-chat-hf",
		"meta-llama/Llama-2-7b-chat-hf",
		"mistralai/Mistral-7B-Instruct-v0.1",
		"mistralai/Mistral-7B-Instruct-v0.2",
		"mistralai/Mixtral-8x7B-Instruct-v0.1",
		"NousResearch/Nous-Capybara-7B-V1p9",
		"NousResearch/Nous-Hermes-2-Mixtral-8x7B-DPO",
		"NousResearch/Nous-Hermes-2-Mixtral-8x7B-SFT",
		"NousResearch/Nous-Hermes-llama-2-7b",
		"NousResearch/Nous-Hermes-Llama2-13b",
		"NousResearch/Nous-Hermes-2-Yi-34B",
		"openchat/openchat-3.5-1210",
		"Open-Orca/Mistral-7B-OpenOrca",
		"togethercomputer/Qwen-7B-Chat",
		"Qwen/Qwen1.5-0.5B-Chat",
		"Qwen/Qwen1.5-1.8B-Chat",
		"Qwen/Qwen1.5-4B-Chat",
		"Qwen/Qwen1.5-7B-Chat",
		"Qwen/Qwen1.5-14B-Chat",
		"Qwen/Qwen1.5-72B-Chat",
		"snorkelai/Snorkel-Mistral-PairRM-DPO",
		"togethercomputer/alpaca-7b",
		"teknium/OpenHermes-2-Mistral-7B",
		"teknium/OpenHermes-2p5-Mistral-7B",
		"togethercomputer/falcon-40b-instruct",
		"togethercomputer/falcon-7b-instruct",
		"togethercomputer/Llama-2-7B-32K-Instruct",
		"togethercomputer/RedPajama-INCITE-Chat-3B-v1",
		"togethercomputer/RedPajama-INCITE-7B-Chat",
		"togethercomputer/StripedHyena-Nous-7B",
		"Undi95/ReMM-SLERP-L2-13B",
		"Undi95/Toppy-M-7B",
		"WizardLM/WizardLM-13B-V1.2",
		"upstage/SOLAR-10.7B-Instruct-v1.0",
	}
}
