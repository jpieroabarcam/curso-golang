package main

import "fmt"

//funcion que se declara dentro de una funcion con retorno
func incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

}
