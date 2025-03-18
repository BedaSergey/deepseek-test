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
	logger logger.Logger
}

func NewPostgresDatabase(logger logger.Logger) *PostgresDatabase {
	if err := godotenv.Load(); err != nil {
		logger.Info(".env file not found, using environment variables")
	}

	logger.Info("Соединяемся с NewPostgresDatabase")
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		logger.Info("Unable to connect to database:", err)
		os.Exit(1)
	}

	err = conn.Ping(context.Background())
	if err != nil {
		logger.Info("Ошибка выполнения функции -conn.Ping-", err)
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

func (db *PostgresDatabase) Close() error {
	return db.Conn.Close(context.Background())
}

func (db *PostgresDatabase) Exec(ctx context.Context, query string, args ...interface{}) (int64, error) {
	result, err := db.Conn.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

func (db *PostgresDatabase) QueryRow(ctx context.Context, query string, args ...interface{}) Row {
	return db.Conn.QueryRow(ctx, query, args...)
}
