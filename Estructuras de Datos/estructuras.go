package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

func main() {
	// var p Persona
	// p.nombre = "Piero"
	// p.edad = 35
	// p.correo = "jabarca@soint.com"

	p := Persona{"Piero", 35, "jabarca@soaint.com"}
	p.edad = 30
	fmt.Println(p)

	p2 := Persona{"juan", 35, "jperez@soaint.com"}
	fmt.Println(p2)

}
