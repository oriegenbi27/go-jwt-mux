package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oriegenbi27/go-jwt-mux/controller/autcontroller"
	"github.com/oriegenbi27/go-jwt-mux/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", autcontroller.Login).Methods("POST")
	r.HandleFunc("/register", autcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", autcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))

}
