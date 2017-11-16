package main

import (
	AuthService "./service/auth"
	"github.com/gorilla/mux"
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
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	///api end point
	api := r.PathPrefix("/api").Subrouter()

	api.HandleFunc("/", welcomeApi)

	//auth Service
	api.HandleFunc("/users", AuthService.Getallusers).Methods("GET", "POST")

	//static html
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./public/")))

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}
