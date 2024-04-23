package main

import "fmt"

// funcion variadica - recibe n cantidad de datos solo un tipo y combinado
func suma(name string, nums ...int) int {
	var total int

	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s la suma es: %d\n", name, total)

	return total
}

//n cantidad de datos y diferentes tipos
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}

}

func main() {
	fmt.Println(suma("Piero", 12, 45, 78, 56))
	fmt.Println(suma("Juan", 10, 20, 30, 40, 50))

	imprimirDatos("Hola", 28, true, 3.14)
}
