package main

import (
	"errors"
	"fmt"
)

func divide(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por Cero") //interfaz de errror
		//eturn 0, fmt.Errorf("No se puede dividir por Cero") // usando fmt
	}
	return dividendo / divisor, nil
}

func main() {
	// str := "123a"
	// num, err := strconv.Atoi(str)
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// fmt.Println("Numero: ", num)

	resultado, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println("Resultado: ", resultado)

}
