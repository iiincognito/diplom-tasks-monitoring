package dbConn

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Path string `envconfig:"TODO_DBFILE" required:"true"`
}

func NewDBConfig() (*DBConfig, error) {
	_ = godotenv.Load(".env")
	var cfg DBConfig
	if err := envconfig.Process("DB", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process env vars: %w", err)
	}
	return &cfg, nil
}
