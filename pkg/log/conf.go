package log

import (
	"github.com/9d77v/band/pkg/env"
)

type Conf struct {
	AppName    string `yaml:"app_name"`
	Level      string `yaml:"log_level"`
	MaxSize    int    `yaml:"log_max_size"`
	MaxBackups int    `yaml:"log_max_backups"`
	MaxAge     int    `yaml:"log_max_age"`
}

func FromEnv() Conf {
	return Conf{
		AppName:    env.String("APP_NAME"),
		Level:      env.String("LOG_LEVEL", LevelDebug),
		MaxSize:    env.Int("LOG_MAX_SIZE", 100),
		MaxBackups: env.Int("LOG_MAX_BACKUPS", 7),
		MaxAge:     env.Int("LOG_MAX_AGE", 30),
	}
}

var (
	LevelDebug = "DEBUG"
	LevelInfo  = "INFO"
	LevelWarn  = "WARN"
	LevelError = "ERROR"
)
