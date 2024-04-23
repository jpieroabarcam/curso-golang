package main

import (
	"fmt"
	"log"

	"github.com/jpieroabarcam/go-greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0) //establecer bandera de formato en cero (sin fecha ni hora)

	names := []string{"Piero", "Juan", "Alex"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// message, err := greetings.Hello("Piero")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	fmt.Println(messages)
}
