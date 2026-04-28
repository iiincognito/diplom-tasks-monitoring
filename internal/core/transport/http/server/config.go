package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" required:"true"`
	ShurdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" required:"true"`
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load(".env")
	var cfg Config
	if err := envconfig.Process("HTTP", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process env vars: %w", err)
	}
	return &cfg, nil
}
