package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

// Config holds all configuration settings
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	Storage  StorageConfig  `yaml:"storage"`
	Auth     AuthConfig     `yaml:"auth"`
}

// ServerConfig defines configuration settings for the web server
type ServerConfig struct {
	Port string `env:"PORT" env-default:"8000"`
}

// DatabaseConfig defines configuration settings for Postgres
type DatabaseConfig struct {
	User     string `env:"POSTGRES_USER" env-required:"true"`
	Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
	Port     string `env:"POSTGRES_PORT" env-required:"true"`
	DBName   string `env:"POSTGRES_DB" env-required:"true"`
	Host     string `env:"POSTGRES_HOST" env-required:"true"`
	SSLMode  string `env:"POSTGRES_SSLMODE" env-required:"true"`
}

// StorageConfig defines configuration settings for MinIO
type StorageConfig struct {
	Endpoint       string `env:"MINIO_ENDPOINT" env-required:"true"`
	SSL            bool   `env:"MINIO_SSL" env-default:"false"`
	BucketName     string `env:"MINIO_BUCKET_NAME" env-required:"true"`
	BucketLocation string `env:"MINIO_BUCKET_LOCATION" env-required:"true"`
}

// AuthConfig holds authentication-related settings
type AuthConfig struct {
	JWTSecret string `env:"JWT_SECRET" env-required:"true"`
}

// Init initializes the application configuration
func Init() (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
