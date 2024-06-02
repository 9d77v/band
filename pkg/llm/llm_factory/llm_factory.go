package llm_factory

import (
	"sync"

	"github.com/9d77v/band/pkg/llm"
	"github.com/9d77v/band/pkg/llm/impl/openai"
	"github.com/9d77v/band/pkg/llm/impl/qwen"
)

var (
	client llm.LLM
	once   sync.Once
)

var (
	TypeOpenAI  = "openai"
	TypeQianWen = "qwen"
)

func NewLLM(conf llm.Conf) (llm.LLM, error) {
	var err error
	var client llm.LLM
	switch conf.Type {
	case TypeQianWen:
		client = qwen.NewQianWen(conf)
	default:
		client = openai.NewOpenAI(conf)
	}
	return client, err
}

func LLMSingleton(conf llm.Conf) (llm.LLM, error) {
	var err error
	once.Do(func() {
		client, err = NewLLM(conf)
	})
	return client, err
}
