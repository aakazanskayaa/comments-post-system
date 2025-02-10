package config

import (
	"os"
)

// Config хранит параметры запуска
type Config struct {
	StorageType string // "memory" или "postgres"
	PostgresDSN string // Строка подключения к PostgreSQL
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() *Config {
	storageType := os.Getenv("STORAGE_TYPE")
	if storageType == "" {
		storageType = "memory" // Значение по умолчанию
	}

	postgresDSN := os.Getenv("POSTGRES_DSN")

	return &Config{
		StorageType: storageType,
		PostgresDSN: postgresDSN,
	}
}
