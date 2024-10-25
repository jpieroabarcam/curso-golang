package main

import (
	"log"
	"net/http"
	"rpsweb/handlers"
)

func main() {
	// crear enrutador
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/game", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servido escuchando en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router)) // un puerto y un manejador de rutas

}
