package main

import (
	"fmt"
	"net/http"

	"alex-j-butler.com/steam-redirect/database"
	"alex-j-butler.com/steam-redirect/models"

	"github.com/codegangsta/martini"
	"github.com/codegangsta/martini-contrib/render"
	"github.com/jinzhu/gorm"
)

func main() {
	m := martini.Classic()
	m.Use(database.Database(database.MySQLConfig{
		Username:  "root",
		Password:  "",
		Database:  "steam_redirect",
		Charset:   "utf8",
		ParseTime: true,
	}, &models.Server{}))
	m.Use(render.Renderer())
	defer database.Close()

	m.Get("/server", ListServers)
	m.Get("/server/:name", HandleServer)

	m.Run()
}

func ListServers(r render.Render, db *gorm.DB) {
	var servers []models.Server

	if err := db.Find(&servers); err.Error == nil {
		r.HTML(200, "list", servers)
	}
}

func HandleServer(w http.ResponseWriter, res *http.Request, params martini.Params, db *gorm.DB) {
	var server models.Server

	if err := db.Where(&models.Server{Name: params["name"]}).First(&server); err.Error == nil {
		url := fmt.Sprintf("steam://connect/%s:%d/%s", server.IP, server.Port, server.Password)
		http.Redirect(w, res, url, http.StatusTemporaryRedirect)
	}
}
