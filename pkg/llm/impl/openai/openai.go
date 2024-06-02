package openai

import (
	"context"
	"errors"
	"io"
	"log"
	"strings"

	"github.com/9d77v/band/pkg/llm"
	ai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	*ai.Client
	conf llm.Conf
}

func NewOpenAI(conf llm.Conf) *OpenAI {
	apiKey := conf.APIKey
	config := ai.DefaultConfig(apiKey)
	baseUrl := strings.TrimSpace(conf.BaseURL)
	if baseUrl != "" {
		config.BaseURL = baseUrl
	}
	log.Println("init openai client:", conf)
	return &OpenAI{
		Client: ai.NewClientWithConfig(config),
		conf:   conf,
	}
}

func (o *OpenAI) GetConf() llm.Conf {
	return o.conf
}

func (o *OpenAI) GenerateImage(ctx context.Context, req *llm.GenerateImageRequest) (*llm.GenerateImageResponse, error) {
	requestBody := ai.ImageRequest{
		Model:          req.Model,
		Quality:        req.Quality,
		Prompt:         req.Prompt,
		Style:          req.Style,
		Size:           req.Size,
		ResponseFormat: ai.CreateImageResponseFormatURL,
		N:              req.N,
	}
	respUrl, err := o.CreateImage(
		ctx,
		requestBody,
	)
	if err != nil {
		log.Printf("Image creation error: %v\n", err)
		return nil, err
	}
	return &llm.GenerateImageResponse{
		ImageUrl:      respUrl.Data[0].URL,
		RevisedPrompt: respUrl.Data[0].RevisedPrompt,
	}, nil
}
func (o *OpenAI) GenerateContentFromImage(ctx context.Context,
	req *llm.GenerateContentFromImageRequest) (*llm.GenerateContentFromImageResponse, error) {
	requestBody := ai.ChatCompletionRequest{
		Model: req.Model,
		Messages: []ai.ChatCompletionMessage{
			{Role: ai.ChatMessageRoleUser, MultiContent: []ai.ChatMessagePart{
				{Type: ai.ChatMessagePartTypeImageURL, ImageURL: &ai.ChatMessageImageURL{
					URL: req.ImageUrl, Detail: ""}},
				{Type: ai.ChatMessagePartTypeText, Text: req.Prompt}}}},
		TopP:           llm.DefaultTopP,
		ResponseFormat: &ai.ChatCompletionResponseFormat{Type: ai.ChatCompletionResponseFormatTypeJSONObject},
	}
	res, err := o.CreateChatCompletion(context.Background(), requestBody)
	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}
	return &llm.GenerateContentFromImageResponse{
		Data: res.Choices[0].Message.Content,
	}, nil
}

func (o *OpenAI) ChatStream(ctx context.Context,
	req *llm.ChatCompletionRequest, respChan chan llm.ChatCompletionStreamResponse) {
	defer close(respChan)
	messages := []ai.ChatCompletionMessage{}
	for _, v := range req.Messages {
		multiContent := []ai.ChatMessagePart{}
		for _, c := range v.MultiContent {
			var imageUrl *ai.ChatMessageImageURL
			if c.ImageURL != nil {
				imageUrl = &ai.ChatMessageImageURL{
					URL:    c.ImageURL.URL,
					Detail: ai.ImageURLDetail(c.ImageURL.Detail),
				}
			}
			multiContent = append(multiContent, ai.ChatMessagePart{
				Type:     ai.ChatMessagePartType(c.Type),
				Text:     c.Text,
				ImageURL: imageUrl,
			})
		}
		if len(multiContent) == 0 {
			multiContent = nil
		}
		messages = append(messages, ai.ChatCompletionMessage{
			Role:         v.Role,
			Content:      v.Content,
			MultiContent: multiContent,
		})
	}
	var responseFormat *ai.ChatCompletionResponseFormat
	if req.ResponseFormat != nil {
		responseFormat = &ai.ChatCompletionResponseFormat{
			Type: ai.ChatCompletionResponseFormatType(req.ResponseFormat.Type),
		}
	}
	requestBody := ai.ChatCompletionRequest{
		Model:          req.Model,
		MaxTokens:      req.MaxTokens,
		Stream:         true,
		Messages:       messages,
		ResponseFormat: responseFormat,
	}
	stream, err := o.CreateChatCompletionStream(ctx, requestBody)
	if err != nil {
		log.Printf("CompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			return
		}
		if err != nil {
			log.Printf("Stream error: %v\n", err)
			return
		}
		choices := []llm.ChatCompletionStreamChoice{}
		for _, v := range response.Choices {
			toolCalls := []llm.ToolCall{}
			for _, tc := range v.Delta.ToolCalls {
				toolCalls = append(toolCalls, llm.ToolCall{
					Index: tc.Index,
					ID:    tc.ID,
					Type:  llm.ToolType(tc.Type),
					Function: llm.FunctionCall{
						Name:      tc.Function.Name,
						Arguments: tc.Function.Arguments,
					},
				})
			}
			if len(toolCalls) == 0 {
				toolCalls = nil
			}
			choices = append(choices, llm.ChatCompletionStreamChoice{
				Index: v.Index,
				Delta: llm.ChatCompletionStreamChoiceDelta{
					Content:   v.Delta.Content,
					Role:      v.Delta.Role,
					ToolCalls: toolCalls,
				},
				FinishReason: llm.FinishReason(v.FinishReason),
			})
		}
		respChan <- llm.ChatCompletionStreamResponse{
			ID:      response.ID,
			Object:  response.Object,
			Created: response.Created,
			Model:   response.Model,
			Choices: choices,
		}
	}
}
