package main

import (
	"fmt"
)

func main() {
	//var a []int
	diasSemana := []string{"domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println(diasSemana)
	diasRebanada := diasSemana[0:5]
	fmt.Println(diasRebanada)

	diasRebanada = append(diasRebanada, "Viernes", "Sabado", "OtroDia")
	fmt.Println(diasRebanada)

	//eliminar elementos
	diasRebanada = append(diasRebanada[:2], diasRebanada[3:]...)
	fmt.Println(diasRebanada)

	fmt.Println(len(diasRebanada))
	fmt.Println(cap(diasRebanada)) // como es una rebanada de otra muestra la capacidad que puede almacenar

	//crear rebanadas a partir de un indice y longitud
	nombres := make([]string, 5, 10)
	nombres[0] = "Alex"
	fmt.Println(nombres)

	//copiar
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	copy(rebanada2, rebanada1)
	fmt.Println(copy(rebanada2, rebanada1))
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)

}
