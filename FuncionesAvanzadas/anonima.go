package main

import "fmt"

// funcion anonima como parametro
func saludar(name string, f func(string)) {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}

func main() {

	//funcion que no tiene nombre - anonima
	saludo := func(name string) {
		fmt.Printf("Hola %s soy una funcion anonima\n", name)
	}
	saludo("Piero")
	saludar("Alex", saludo)

	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 5)

	fmt.Println(r1, r2)

}
