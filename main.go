package main

import (
	"fmt"
	"net/http"

	"alex-j-butler.com/go-redirect/database"

	"github.com/codegangsta/martini"
	"github.com/jinzhu/gorm"
)

func main() {
	m := martini.Classic()
	m.Use(database.Database())
	defer database.Close()

	database.DB.AutoMigrate(&Server{})

	m.Get("/server/:name", HandleServer)

	m.Run()
}

func HandleServer(w http.ResponseWriter, r *http.Request, params martini.Params, db *gorm.DB) {
	var server *Server
	db.Where(Server{Name: params["name"]}).First(&server)

	url := fmt.Sprintf("steam://connect/%s:%d/%s", server.IP, server.Port, server.Password)
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}
