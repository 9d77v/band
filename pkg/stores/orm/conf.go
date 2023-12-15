package orm

import "github.com/9d77v/band/pkg/env"

type Conf struct {
	Driver       string `yaml:"driver"`
	Host         string `yaml:"host"`
	Port         uint   `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	DBName       string `yaml:"dbname"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	TablePrefix  string `yaml:"table_prefix"`
	Debug        bool   `yaml:"debug"`
}

func FromEnv() Conf {
	return Conf{
		Driver:       env.String("DB_DRIVER", "postgres"),
		Host:         env.String("DB_HOST", "127.0.0.1"),
		Port:         uint(env.Int("DB_PORT", 5432)),
		User:         env.String("DB_USER", "postgres"),
		Password:     env.String("DB_PASSWORD", "123456"),
		DBName:       env.String("DB_NAME"),
		MaxIdleConns: env.Int("DB_MAX_IDEL_CONNS", 100),
		MaxOpenConns: env.Int("DB_OPEN_IDEL_CONNS", 100),
		TablePrefix:  env.String("DB_TABLE_PREFIX"),
		Debug:        env.Bool("DEBUG", true),
	}
}
