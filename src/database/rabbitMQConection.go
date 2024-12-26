package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	RabbitMQURL string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Erro ao carregar arquivo .env: %s", err)
	}

	return &Config{
		RabbitMQURL: os.Getenv("RABBITMQ_URL"),
	}
}
