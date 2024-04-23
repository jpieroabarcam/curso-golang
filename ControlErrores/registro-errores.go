package main

import (
	"log"
	"os"
)

func main() {
	// log.Fatal("Oye soy registro de errores")
	// log.Print("puedes leerme?")

	// log.Panic("Oye soy registro de errores")
	// log.Print("puedes leerme?")

	// log.SetPrefix("main():") // agrega un prefijo a los logs

	// log.Print("Este es un mensaje de registro")
	// log.Println("Este es otro mensaje de registro")
	log.SetPrefix("main():")
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // para registro de los errores

	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
	log.Print("¡Oye soy un Log!")

	defer file.Close()

}
