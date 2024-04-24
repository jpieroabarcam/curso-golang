# Saludos en Go
Este paquete proporciona una forma simple de obtener saludos de persona

## Instalacion
Ejecuta el siguiente compando para instalar el paguete

```bash
go get -u github.com/jpieroabarcam/go-greetings
```

## Uso
Aqui tienes un ejemplo de como utilizar el paquete en tu codigo

```go
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
```

Este ejemplo importa el paquete del repositorio y llama a la funcion Hello con un saludo personalizado. Si ocurre un error, se imprime un mensaje de error.