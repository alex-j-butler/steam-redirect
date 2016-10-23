package main

import (
	"fmt"
	"net/http"

	"alex-j-butler.com/steam-redirect/database"
	"alex-j-butler.com/steam-redirect/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func main() {
	m := mux.NewRouter()

	DB = database.Dial(database.MySQLConfig{
		Username:  "root",
		Password:  "",
		Database:  "steam_redirect",
		Charset:   "utf8",
		ParseTime: true,
	}, &models.Server{})
	defer DB.Close()

	m.HandleFunc("/server/{name}", ServerHandle).
		Methods("GET")

	http.ListenAndServe(":3000", m)
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
