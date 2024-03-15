package neo4j

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func FromEnv() Conf {
	return Conf{
		Host:     env.String("NEO4j_HOST", "127.0.0.1"),
		Port:     uint(env.Int("NEO4j_PORT", 7687)),
		Username: env.String("NEO4j_USERNAME", "neo4j"),
		Password: env.String("NEO4j_PASSWORD", "12345678"),
	}
}
