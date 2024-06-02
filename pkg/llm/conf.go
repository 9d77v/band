package llm

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Type    string `yaml:"type"`
	APIKey  string `yaml:"api_key"`
	Model   string `yaml:"model"`
	BaseURL string `yaml:"base_url"`
}

func FromEnv() Conf {
	return Conf{
		Type:    env.String("LLM_TYPE"),
		APIKey:  env.String("LLM_API_KEY"),
		Model:   env.String("LLM_MODEL"),
		BaseURL: env.String("LLM_BASE_URL"),
	}
}
