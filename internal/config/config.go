package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	// Config holds various configuration settings
	Config struct {
		Server   HTTPConfig
		Postgres PostgresConfig
		Minio    MinioConfig
	}

	// HTTPConfig defines configuration settings for web server
	HTTPConfig struct {
		Port string `env:"PORT" env-default:"8000"`
	}

	// PostgresConfig defines configuration settings for PostgresConfig
	PostgresConfig struct {
		User     string `env:"POSTGRES_USER"     env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		Port     string `env:"POSTGRES_PORT"     env-required:"true"`
		DBName   string `env:"POSTGRES_DB"       env-required:"true"`
		Host     string `env:"POSTGRES_HOST"     env-required:"true"`
		SSLMode  string `env:"POSTGRES_SSLMODE"  env-required:"true"`
	}

	// MinioConfig defines configuration settings for Minio
	MinioConfig struct {
		Endpoint       string `env:"MINIO_ENDPOINT"        env-required:"true"`
		SSL            string `env:"MINIO_SSL_MODE"        env-required:"true"`
		BucketName     string `env:"MINIO_BUCKET_NAME"     env-required:"true"`
		BucketLocation string `env:"MINIO_BUCKET_LOCATION" env-required:"true"`
	}
)

// Init initialize app configuration
func Init() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
