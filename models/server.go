package models

import "github.com/jinzhu/gorm"

type Server struct {
	gorm.Model
	Name     string
	IP       string
	Port     int
	Password string
}
