package main

import (
	Config "./config"
	AuthService "./service/auth"
	ModelUsers "./service/auth/model"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func welcome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to blog Service"))
}

func welcomeApi(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome to blog Api"))
}

func main() {
	//migrate database
	config := Config.GetConfig()

	dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
		config.DB.Username,
		config.DB.Password,
		config.DB.Name,
		config.DB.Charset)

	db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatal("Could not connect database")
	}

	ModelUsers.DBMigrate(db)

	//builth router
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	///api end point
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", welcomeApi)

	//auth Service
	api.HandleFunc("/users", AuthService.Getallusers).Methods("GET")

	api.HandleFunc("/users", AuthService.CreateUsers).Methods("POST")

	//static html
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}
