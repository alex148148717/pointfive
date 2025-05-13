package config

import (
	"fmt"
	"os"
)

type Config struct {
	DatabaseURL        string
	ServerPort         string
	TaskQueueImportJob string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DatabaseURL:        getEnv("DSN", "host=localhost user=postgres password=postgres dbname=dbt port=5432 sslmode=disable TimeZone=UTC"),
		ServerPort:         getEnv("PORT", "8081"),
		TaskQueueImportJob: getEnv("TASK_QUEUE_IMPORT_JOB", "TASK_QUEUE_IMPORT_JOB"),
	}

	if cfg.DatabaseURL == "" {
		return nil, fmt.Errorf("missing required environment variable: DATABASE_URL")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
