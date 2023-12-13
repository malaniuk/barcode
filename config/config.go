package config

import (
	"barcode/packages/env"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	TgApiKey     string
	TgAllowedIds []int64
}

func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	return &Config{
		TgApiKey:     env.GetString("TG_API_KEY"),
		TgAllowedIds: env.GetInt64Array("TG_ALLOWED_IDS"),
	}
}
