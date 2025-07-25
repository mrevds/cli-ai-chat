package config

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey string
}

func LoadConfig() (*Config, error) {
	// В реальном приложении лучше использовать файл конфигурации или переменные окружения
	apiKey := os.Getenv("DEEPSEEK_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("необходимо установить переменную окружения DEEPSEEK_API_KEY")
	}

	return &Config{
		APIKey: apiKey,
	}, nil
}
