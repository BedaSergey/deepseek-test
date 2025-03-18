package logger

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	logger *logrus.Logger
}

func NewLogrusLogger() *LogrusLogger {
	logger := logrus.New()

	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Ошибка при загрузке .env файла:", err)
	}
	logrus.SetLevel(logrus.InfoLevel)
	file, err := os.OpenFile(os.Getenv("LOG_DIR"), os.O_RDWR, 0666)
	if err != nil {
		logrus.Fatal("Не удалось открыть файл для логирования:", err)
	}
	logger.SetOutput(file)

	return &LogrusLogger{logger: logger}
}

func (l *LogrusLogger) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *LogrusLogger) Error(args ...interface{}) {
	l.logger.Error(args...)
}
