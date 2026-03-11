package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host        string        `env:"HTTP_HOST" env-default:"0.0.0.0"`
	Port        string        `env:"HTTP_PORT" env-default:"8080"`
	Timeout     time.Duration `env:"HTTP_TIMEOUT" env-default:"5s"`
	IdleTimeout time.Duration `env:"HTTP_IDLE_TIMEOUT" env-default:"60s"`
}

func Get() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
