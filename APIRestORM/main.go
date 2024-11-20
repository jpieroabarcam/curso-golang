package main

import (
	"gorm/handlers"
	"gorm/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	models.MigrarUser()

	// Rutas
	muxs := mux.NewRouter()

	// EndPoint
	muxs.HandleFunc("/api/user/", handlers.GetUsers).Methods("GET")
	muxs.HandleFunc("/api/user/{id:[0-9]+}", handlers.GetUser).Methods("GET")
	muxs.HandleFunc("/api/user/", handlers.CreateUser).Methods("POST")
	muxs.HandleFunc("/api/user/{id:[0-9]+}", handlers.UpdateUser).Methods("PUT")
	muxs.HandleFunc("/api/user/{id:[0-9]+}", handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", muxs))
}
