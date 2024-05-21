package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL   string
	LISTEN_ADDRESS string
}

func Load() (*Config, error) {
	slog.Debug("loading .env file")
	if godotenv.Load() != nil {
		slog.Warn("failed to load .env file")
	}

	DATABASE_URL, err := env("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	LISTEN_ADDRESS, err := env("LISTEN_ADDRESS")
	if err != nil {
		return nil, err
	}

	config := Config{
		DATABASE_URL,
		LISTEN_ADDRESS,
	}

	return &config, nil
}

func env(name string) (string, error) {
	value, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf("missing env: %s", name)
	}
	return value, nil
}
