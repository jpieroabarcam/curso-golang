package main

import "fmt"

// si no se implementa un panic sin necesidad de implementarlo
func dividir(dividendo, divisor int) {
	defer func() { // funcion anonima de capturar cualquier panico
		// ya no se interrumpe la ejecucion de programa
		if r := recover(); r != nil { //recover: recuperar ese panico
			fmt.Println(r)
		}
	}()
	validateZero(divisor)
	fmt.Println(dividendo / divisor)
}

func validateZero(divisor int) {
	if divisor == 0 {
		panic("No se puede dividir por cero") // provocar un panico en la aplicacion
	}
}

func main() {
	dividir(100, 10)
	dividir(200, 25)
	dividir(34, 0)
	dividir(100, 4)

}
