package config

import (
	"log"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL   string
	LISTEN_ADDRESS string
}

func Load() Config {
	if godotenv.Load() != nil {
		slog.Warn("failed to load .env")
	}

	config := Config{
		DATABASE_URL:   env("DATABASE_URL"),
		LISTEN_ADDRESS: env("LISTEN_ADDRESS"),
	}

	return config
}

func env(name string) string {
	value, ok := os.LookupEnv(name)
	if !ok {
		log.Fatalf("missing env: %s\n", name)
	}
	return value
}
