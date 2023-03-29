package config

import (
	"os"
)

type Config struct {
	DBName   string
	LogLevel string
}

func NewConfig() *Config {
	return &Config{
		DBName:   getEnv("DB_NAME", "smartassist"),
		LogLevel: getEnv("LOG_LEVEL", "info"),
	}
}

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
