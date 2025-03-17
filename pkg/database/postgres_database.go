package database

import (
	"context"
	"os"
	"rent_alice/pkg/logger"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type PostgresDatabase struct {
	Conn   *pgx.Conn
	logger *logger.LogrusLogger
}

func NewPostgresDatabase(logger *logger.LogrusLogger) *PostgresDatabase {
	err := godotenv.Load()
	if err != nil {
		logger.Error("Ошибка при загрузке .env файла:", err)
	}

	logger.Info("Соединяемся с NewPostgresDatabase")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Error("Unable to connect to database:", err)
		os.Exit(1)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		logger.Error("Ошибка выполнения функции -conn.Ping-", err)
	}

	return &PostgresDatabase{Conn: conn, logger: logger}
}

func (db *PostgresDatabase) PingConnection() error {
	err := db.Conn.Ping(context.Background())
	if err != nil {
		// db.logger.Error("Ошибка функции -conn.Ping-:", err)
		return err
	}
	return nil
}
