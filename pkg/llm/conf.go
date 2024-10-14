package llm

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type    string `yaml:"type"`
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
}

func FromEnv() Conf {
	return Conf{
		Type:    env.String("LLM_TYPE"),
		APIKey:  env.String("LLM_API_KEY"),
		BaseURL: env.String("LLM_BASE_URL"),
	}
}
