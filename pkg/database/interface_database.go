package database

type Database interface {
	PingConnection() error
}
