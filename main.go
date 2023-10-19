package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oriegenbi27/go-jwt-mux/controller/autcontroller"
	"github.com/oriegenbi27/go-jwt-mux/controller/footballcontroller"
	"github.com/oriegenbi27/go-jwt-mux/controller/productcontroller"
	"github.com/oriegenbi27/go-jwt-mux/middlewares"
	"github.com/oriegenbi27/go-jwt-mux/models"
)

func main() {

	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", autcontroller.Login).Methods("POST")
	r.HandleFunc("/register", autcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", autcontroller.Logout).Methods("GET")

	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productcontroller.Index).Methods("GET")
	api.HandleFunc("/news", footballcontroller.News).Methods("GET")

	api.Use(middlewares.JWTMidlleware)

	log.Fatal(http.ListenAndServe(":8080", r))

}
