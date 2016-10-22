package database

import (
	"github.com/codegangsta/martini"
	"github.com/jinzhu/gorm"

	// MySQL dialect for GORM
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Database(config DatabaseConfig, values ...interface{}) martini.Handler {
	// Connect to gorm using the DatabaseConfig.
	DB, err := gorm.Open(config.getDialect(), config.getConnectString())

	if err != nil {
		panic(err)
	}

	// Auto migrate the structs provided.
	DB.AutoMigrate(values...)

	return func(c martini.Context) {
		database := DB
		c.Map(database)
		c.Next()
	}
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
