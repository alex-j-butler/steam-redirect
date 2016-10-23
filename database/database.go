package database

import (
	"github.com/jinzhu/gorm"
)

func Dial(config DatabaseConfig, values ...interface{}) *gorm.DB {
	// Connect to gorm using the DatabaseConfig.
	DB, err := gorm.Open(config.getDialect(), config.getConnectString())

	if err != nil {
		panic(err)
	}

	// Auto migrate the structs provided.
	DB.AutoMigrate(values...)

	return DB
}
