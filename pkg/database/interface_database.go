package database

import "context"

type Database interface {
	PingConnection() error
	Exec(ctx context.Context, query string, args ...interface{}) (int64, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) Row
}

type Row interface {
	Scan(dest ...interface{}) error
}
