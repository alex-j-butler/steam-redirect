package database

import (
	"github.com/codegangsta/martini"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func Database() martini.Handler {
	DB, err := gorm.Open("sqlite", "test.sqlite")

	if err != nil {
		panic(err)
	}

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
