package main

import (
	"fmt"
	"net/http"

	"alex-j-butler.com/steam-redirect/config"
	"alex-j-butler.com/steam-redirect/database"
	"alex-j-butler.com/steam-redirect/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	config.InitialiseConfiguration()

	m := mux.NewRouter()

	DB = database.Dial(database.MySQLConfig{
		Username:  config.Conf.Database.Username,
		Password:  config.Conf.Database.Password,
		Database:  config.Conf.Database.Database,
		Charset:   config.Conf.Database.Charset,
		ParseTime: true,
	}, &models.Server{})
	defer DB.Close()

	m.HandleFunc("/server/{name}", ServerHandle).
		Methods("GET")

	m.NotFoundHandler = http.HandlerFunc(NotFoundHandle)

	http.ListenAndServe(fmt.Sprintf("%s:%d", config.Conf.Address, config.Conf.Port), m)
}

func ServerHandle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]

	var server models.Server
	if err := DB.Where(&models.Server{Name: name}).First(&server); err.Error == nil {
		url := fmt.Sprintf("steam://connect/%s:%d/%s", server.IP, server.Port, server.Password)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	} else {
		w.Write([]byte("Server does not exist."))
	}
}

func NotFoundHandle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server does not exist."))
}
