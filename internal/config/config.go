package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server   HTTPConfig
		Postgres PostgresConfig
	}

	HTTPConfig struct {
		Port string `env:"PORT" env-default:"8000"`
	}

	PostgresConfig struct {
		User     string `env:"POSTGRES_USER"     env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Port     string `env:"POSTGRES_PORT"     env-required:"true"`
		DBName   string `env:"POSTGRES_DB"       env-required:"true"`
		Host     string `env:"POSTGRES_HOST"     env-required:"true"`
		SSLMode  string `env:"POSTGRES_SSLMODE"  env-required:"true"`
	}
)

func Init() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
