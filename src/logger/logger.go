package logs

import (
	"io"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupLogger() {

	if err := godotenv.Load(); err != nil {
		logrus.Warn("Erro ao carregar o arquivo .env, usando configurações padrão")
	}

	logFilePath := os.Getenv("LOG_FILE_PATH")
	maxSize, _ := strconv.Atoi(os.Getenv("LOG_MAX_SIZE"))
	maxBackups, _ := strconv.Atoi(os.Getenv("LOG_MAX_BACKUPS"))
	maxAge, _ := strconv.Atoi(os.Getenv("LOG_MAX_AGE"))

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	fileLogger := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   true,
	}

	multiWriter := io.MultiWriter(fileLogger, logrus.StandardLogger().Out)

	logrus.SetOutput(multiWriter)
	logrus.SetLevel(logrus.InfoLevel)
}
