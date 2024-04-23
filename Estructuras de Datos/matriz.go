package main

import "fmt"

func main() {

	var a = [...]int{200,100,1,2,3,4,5}

	for index, value := range a {
		fmt.Printf("los valores de %d es %d\n", index, value)
	}

	var matriz = [3][3] int{{1,2,3}, {4,5,6}, {7,8,9}}

	fmt.Println(matriz)


}