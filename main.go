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

type App struct {
	DB *gorm.DB
}

func welcome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Welcome to blog Service"))
}

func welcomeApi(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("welcome to blog Api"))
}

func (a *App) aplication() {
	Migrate()
	Router()

}

func Migrate() {
	//get connecttion db
	db := Config.Db()

	ModelUsers.DBMigrate(db)
}

func Router() {
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
	fmt.Println("server listen in port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}

func main() {

	app := App{}
	app.aplication()
}
