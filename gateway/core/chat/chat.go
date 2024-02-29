package chat

type ResponseFormat struct {
	Type string `json:"type,omitempty"`
}

type ChatCompletionRequest struct {
	Model            string                  `json:"model"`
	Messages         []ChatCompletionMessage `json:"messages"`
	Temperature      float32                 `json:"temperature,omitempty" default:"1"`
	Suffix           string                  `json:"suffix,omitempty"`
	Seed             uint64                  `json:"seed,omitempty"`
	N                uint64                  `json:"n,omitempty"`
	Echo             bool                    `json:"echo,omitempty"`
	BestOf           uint64                  `json:"best_of,omitempty"`
	PresencePenalty  float32                 `json:"presence_penalty,omitempty"`
	FrequencyPenalty float32                 `json:"frequency_penalty,omitempty"`
	Stream           bool                    `json:"stream,omitempty"`
	TopK             float32                 `json:"top_k,omitempty"`
	TopP             float32                 `json:"top_p,omitempty"`
	Stop             []string                `json:"stop,omitempty"`
	MaxTokens        uint64                  `json:"max_tokens,omitempty"`
	LogProbs         bool                    `json:"logprobs,omitempty"`
	TopLogprobs      uint64                  `json:"top_logprobs,omitempty"`
	LogitBias        map[string]any          `json:"logit_bias,omitempty"`
	ToolChoice       map[string]any          `json:"tool_choice,omitempty"`
	User             string                  `json:"user,omitempty"`
}

type Usage struct {
	PromptTokens     int64 `json:"prompt_tokens,omitempty"`
	CompletionTokens int64 `json:"completion_tokens,omitempty"`
	TotalTokens      int64 `json:"total_tokens,omitempty"`
}

type LogprobResult struct {
	Tokens        []string             `json:"tokens"`
	TokenLogprobs []float32            `json:"token_logprobs"`
	TopLogprobs   []map[string]float32 `json:"top_logprobs"`
	TextOffset    []int                `json:"text_offset"`
}

type Function struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type ToolCallMessage struct {
	Id       string   `json:"id"`
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type ChatCompletionMessage struct {
	Role      string            `json:"role"`
	Content   string            `json:"content"`
	LogProbs  map[string]any    `json:"logprobs,omitempty"`
	ToolCalls []ToolCallMessage `json:"tool_calls,omitempty"`
}

type CompletionChoice struct {
	Index        int                   `json:"index"`
	Message      ChatCompletionMessage `json:"message"`
	FinishReason string                `json:"finish_reason"`
	LogProbs     LogprobResult         `json:"logprobs"`
}

type ChatCompletionResponse struct {
	ID                string             `json:"id,omitempty"`
	Object            string             `json:"object,omitempty"`
	Created           uint64             `json:"created,omitempty"`
	Model             string             `json:"model,omitempty"`
	Choices           []CompletionChoice `json:"choices,omitempty"`
	Usage             Usage              `json:"usage,omitempty"`
	Cached            bool               `json:"cached,omitempty" default:"false"`
	SystemFingerprint string             `json:"system_fingerprint,omitempty"`
}
