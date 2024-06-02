package qwen

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/9d77v/band/pkg/llm"
	"github.com/9d77v/band/pkg/utils"
	"github.com/joho/godotenv"
)

var llmClient *QianWen

func setup() {
	fmt.Println("Before all tests")
	if utils.FileExist(".env") {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}
	}
	llmConf := llm.FromEnv()
	llmClient = NewQianWen(llmConf)
}

func teardown() {
	fmt.Println("After all tests")
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestQianWen_ChatStream(t *testing.T) {
	type args struct {
		ctx      context.Context
		req      *llm.ChatCompletionRequest
		respChan chan llm.ChatCompletionStreamResponse
	}
	tests := []struct {
		name string
		args args
	}{
		{"test chat stream", args{context.Background(),
			&llm.ChatCompletionRequest{
				Model: "qwen-max",
				Messages: []llm.ChatCompletionMessage{
					{
						Role:    "user",
						Content: "你是谁",
					},
				},
			}, make(chan llm.ChatCompletionStreamResponse, 1)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			go llmClient.ChatStream(tt.args.ctx, tt.args.req, tt.args.respChan)
			for c := range tt.args.respChan {
				fmt.Print(c.Choices[0].Delta.Content)
			}
		})
	}
	t.Fail()
}
