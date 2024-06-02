package llm

import (
	"context"
	"encoding/json"
	"errors"
)

type LLM interface {
	GetConf() Conf
	GenerateImage(ctx context.Context, req *GenerateImageRequest) (*GenerateImageResponse, error)
	GenerateContentFromImage(ctx context.Context, req *GenerateContentFromImageRequest) (*GenerateContentFromImageResponse, error)
	ChatStream(ctx context.Context, req *ChatCompletionRequest, respChan chan ChatCompletionStreamResponse)
}

const DefaultTopP = 0.8

type GenerateImageRequest struct {
	Model          string
	Quality        string
	Prompt         string
	NegativePrompt string
	Size           string
	N              int
	RefImg         string
	RefMode        string
	Style          string
	Seed           int
}
type GenerateImageResponse struct {
	RevisedPrompt string
	ImageUrl      string
}

type GenerateContentFromImageRequest struct {
	Model    string
	Prompt   string
	ImageUrl string
}
type GenerateContentFromImageResponse struct {
	Data string
}

type ChatCompletionResponseFormatType string

const (
	ChatCompletionResponseFormatTypeJSONObject ChatCompletionResponseFormatType = "json_object"
	ChatCompletionResponseFormatTypeText       ChatCompletionResponseFormatType = "text"
)

type ChatCompletionResponseFormat struct {
	Type ChatCompletionResponseFormatType `json:"type,omitempty"`
}

// ChatCompletionRequest represents a request structure for chat completion API.
type ChatCompletionRequest struct {
	Model            string                        `json:"model"`
	Messages         []ChatCompletionMessage       `json:"messages"`
	MaxTokens        int                           `json:"max_tokens,omitempty"`
	Temperature      float32                       `json:"temperature,omitempty"`
	TopP             float32                       `json:"top_p,omitempty"`
	N                int                           `json:"n,omitempty"`
	Stop             []string                      `json:"stop,omitempty"`
	PresencePenalty  float32                       `json:"presence_penalty,omitempty"`
	ResponseFormat   *ChatCompletionResponseFormat `json:"response_format,omitempty"`
	Seed             *int                          `json:"seed,omitempty"`
	FrequencyPenalty float32                       `json:"frequency_penalty,omitempty"`
	User             string                        `json:"user,omitempty"`
	Tools            []Tool                        `json:"tools,omitempty"`
	ToolChoice       any                           `json:"tool_choice,omitempty"`
}

type ChatMessagePartType string

const (
	ChatMessagePartTypeText     ChatMessagePartType = "text"
	ChatMessagePartTypeImageURL ChatMessagePartType = "image_url"
)

type ChatMessagePart struct {
	Type     ChatMessagePartType  `json:"type,omitempty"`
	Text     string               `json:"text,omitempty"`
	ImageURL *ChatMessageImageURL `json:"image_url,omitempty"`
}

type ImageURLDetail string

const (
	ImageURLDetailHigh ImageURLDetail = "high"
	ImageURLDetailLow  ImageURLDetail = "low"
	ImageURLDetailAuto ImageURLDetail = "auto"
)

type ChatMessageImageURL struct {
	URL    string         `json:"url,omitempty"`
	Detail ImageURLDetail `json:"detail,omitempty"`
}

type ChatCompletionMessage struct {
	Role         string `json:"role"`
	Content      string `json:"content"`
	MultiContent []ChatMessagePart

	// This property isn't in the official documentation, but it's in
	// the documentation for the official library for python:
	// - https://github.com/openai/openai-python/blob/main/chatml.md
	// - https://github.com/openai/openai-cookbook/blob/main/examples/How_to_count_tokens_with_tiktoken.ipynb
	Name string `json:"name,omitempty"`

	FunctionCall *FunctionCall `json:"function_call,omitempty"`

	// For Role=assistant prompts this may be set to the tool calls generated by the model, such as function calls.
	ToolCalls []ToolCall `json:"tool_calls,omitempty"`

	// For Role=tool prompts this should be set to the ID given in the assistant's prior request to call a tool.
	ToolCallID string `json:"tool_call_id,omitempty"`
}

var (
	ErrChatCompletionInvalidModel       = errors.New("this model is not supported with this method, please use CreateCompletion client method instead") //nolint:lll
	ErrChatCompletionStreamNotSupported = errors.New("streaming is not supported with this method, please use CreateChatCompletionStream")              //nolint:lll
	ErrContentFieldsMisused             = errors.New("can't use both Content and MultiContent properties simultaneously")
)

func (m ChatCompletionMessage) MarshalJSON() ([]byte, error) {
	if m.Content != "" && m.MultiContent != nil {
		return nil, ErrContentFieldsMisused
	}
	if len(m.MultiContent) > 0 {
		msg := struct {
			Role         string            `json:"role"`
			Content      string            `json:"-"`
			MultiContent []ChatMessagePart `json:"content,omitempty"`
			Name         string            `json:"name,omitempty"`
			FunctionCall *FunctionCall     `json:"function_call,omitempty"`
			ToolCalls    []ToolCall        `json:"tool_calls,omitempty"`
			ToolCallID   string            `json:"tool_call_id,omitempty"`
		}(m)
		return json.Marshal(msg)
	}
	msg := struct {
		Role         string            `json:"role"`
		Content      string            `json:"content"`
		MultiContent []ChatMessagePart `json:"-"`
		Name         string            `json:"name,omitempty"`
		FunctionCall *FunctionCall     `json:"function_call,omitempty"`
		ToolCalls    []ToolCall        `json:"tool_calls,omitempty"`
		ToolCallID   string            `json:"tool_call_id,omitempty"`
	}(m)
	return json.Marshal(msg)
}

func (m *ChatCompletionMessage) UnmarshalJSON(bs []byte) error {
	msg := struct {
		Role         string `json:"role"`
		Content      string `json:"content"`
		MultiContent []ChatMessagePart
		Name         string        `json:"name,omitempty"`
		FunctionCall *FunctionCall `json:"function_call,omitempty"`
		ToolCalls    []ToolCall    `json:"tool_calls,omitempty"`
		ToolCallID   string        `json:"tool_call_id,omitempty"`
	}{}
	if err := json.Unmarshal(bs, &msg); err == nil {
		*m = ChatCompletionMessage(msg)
		return nil
	}
	multiMsg := struct {
		Role         string `json:"role"`
		Content      string
		MultiContent []ChatMessagePart `json:"content"`
		Name         string            `json:"name,omitempty"`
		FunctionCall *FunctionCall     `json:"function_call,omitempty"`
		ToolCalls    []ToolCall        `json:"tool_calls,omitempty"`
		ToolCallID   string            `json:"tool_call_id,omitempty"`
	}{}
	if err := json.Unmarshal(bs, &multiMsg); err != nil {
		return err
	}
	*m = ChatCompletionMessage(multiMsg)
	return nil
}

type ToolCall struct {
	// Index is not nil only in chat completion chunk object
	Index    *int         `json:"index,omitempty"`
	ID       string       `json:"id"`
	Type     ToolType     `json:"type"`
	Function FunctionCall `json:"function"`
}

type FunctionCall struct {
	Name string `json:"name,omitempty"`
	// call function with arguments in JSON format
	Arguments string `json:"arguments,omitempty"`
}

type ToolType string

const (
	ToolTypeFunction ToolType = "function"
)

type Tool struct {
	Type     ToolType            `json:"type"`
	Function *FunctionDefinition `json:"function,omitempty"`
}

type ToolChoice struct {
	Type     ToolType     `json:"type"`
	Function ToolFunction `json:"function,omitempty"`
}

type ToolFunction struct {
	Name string `json:"name"`
}

type FunctionDefinition struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Parameters  any    `json:"parameters"`
}

// ChatCompletionResponse represents a response structure for chat completion API.
type ChatCompletionResponse struct {
	ID      string                 `json:"id"`
	Object  string                 `json:"object"`
	Created int64                  `json:"created"`
	Model   string                 `json:"model"`
	Choices []ChatCompletionChoice `json:"choices"`
}

type ChatCompletionChoice struct {
	Index   int                   `json:"index"`
	Message ChatCompletionMessage `json:"message"`
	// FinishReason
	// stop: API returned complete message,
	// or a message terminated by one of the stop sequences provided via the stop parameter
	// length: Incomplete model output due to max_tokens parameter or token limit
	// function_call: The model decided to call a function
	// content_filter: Omitted content due to a flag from our content filters
	// null: API response still in progress or incomplete
	FinishReason FinishReason `json:"finish_reason"`
	LogProbs     *LogProbs    `json:"logprobs,omitempty"`
}

type TopLogProbs struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes,omitempty"`
}

// LogProb represents the probability information for a token.
type LogProb struct {
	Token   string  `json:"token"`
	LogProb float64 `json:"logprob"`
	Bytes   []byte  `json:"bytes,omitempty"` // Omitting the field if it is null
	// TopLogProbs is a list of the most likely tokens and their log probability, at this token position.
	// In rare cases, there may be fewer than the number of requested top_logprobs returned.
	TopLogProbs []TopLogProbs `json:"top_logprobs"`
}

// LogProbs is the top-level structure containing the log probability information.
type LogProbs struct {
	// Content is a list of message content tokens with log probability information.
	Content []LogProb `json:"content"`
}

type FinishReason string

const (
	FinishReasonStop          FinishReason = "stop"
	FinishReasonLength        FinishReason = "length"
	FinishReasonFunctionCall  FinishReason = "function_call"
	FinishReasonToolCalls     FinishReason = "tool_calls"
	FinishReasonContentFilter FinishReason = "content_filter"
	FinishReasonNull          FinishReason = "null"
)

func (r FinishReason) MarshalJSON() ([]byte, error) {
	if r == FinishReasonNull || r == "" {
		return []byte("null"), nil
	}
	return []byte(`"` + string(r) + `"`), nil
}

type ChatCompletionStreamChoiceDelta struct {
	Content      string        `json:"content,omitempty"`
	Role         string        `json:"role,omitempty"`
	FunctionCall *FunctionCall `json:"function_call,omitempty"`
	ToolCalls    []ToolCall    `json:"tool_calls,omitempty"`
}

type ChatCompletionStreamChoice struct {
	Index        int                             `json:"index"`
	Delta        ChatCompletionStreamChoiceDelta `json:"delta"`
	FinishReason FinishReason                    `json:"finish_reason"`
}

type ChatCompletionStreamResponse struct {
	ID      string                       `json:"id"`
	Object  string                       `json:"object"`
	Created int64                        `json:"created"`
	Model   string                       `json:"model"`
	Choices []ChatCompletionStreamChoice `json:"choices"`
}
