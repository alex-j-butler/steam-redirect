package database

type DatabaseConfig interface {
	getConnectString() string
	getDialect() string
}
