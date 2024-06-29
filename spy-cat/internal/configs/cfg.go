package configs

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"spy-cat/pkg/validator"
)

type Config struct {
	DbHost     string `env:"DB_HOST" validate:"required"`
	DbPort     string `env:"DB_PORT" validate:"required"`
	DbUser     string `env:"DB_USER" validate:"required"`
	DbPassword string `env:"DB_PASSWORD" validate:"required"`
	DbName     string `env:"DB_NAME" validate:"required"`

	Api struct {
		Port int `env:"API_PORT" validate:"required"`
	}
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("unable to load ..env file: %w", err)
	}

	cfg := &Config{}
	err = env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ennvironment variables: %w", err)
	}

	if err = validator.GetValidator().Validate(cfg); err != nil {
		return nil, fmt.Errorf("invalid configs values: %w", err)
	}

	return cfg, nil
}
