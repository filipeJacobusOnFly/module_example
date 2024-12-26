package database

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	RabbitMQURL string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Erro ao carregar arquivo .env: %s", err)
	}

	return &Config{
		RabbitMQURL: os.Getenv("RABBITMQ_URL"),
	}
}
