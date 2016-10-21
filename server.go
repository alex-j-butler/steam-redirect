package main

type Server struct {
    gorm.Model
    IP string
    Port int
    Password string
}
