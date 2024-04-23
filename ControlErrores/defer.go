package main

import (
	"fmt"
	"os"
)

func main() {
	// defer fmt.Println(3) // posponer la ejecucion, se ejecuta al final
	// defer fmt.Println(2)
	// defer fmt.Println(1)
	// // cuando todas tienen defer, se agregan a una pila, LIFO

	file, err := os.Create("hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close() // posponer la ejecucion antes que termine main

	_, err = file.Write([]byte("Hola Amigo mundo"))
	if err != nil {
		fmt.Println(err)
		//file.Close() // llamar 2 veces
		return
	}
	//file.Close()
}
