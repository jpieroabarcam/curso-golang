package main

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// pertenece a la structura por el puntero
func (p *Persona) diHola() {
	fmt.Println("Hola, mi nombre es", p.nombre)
}

func main() {

	// var x int = 10
	// var p *int = &x //puntero a la referencia de la memoria de x a p
	// fmt.Println(&x)
	// fmt.Println(p)

	// fmt.Println(x)
	// editar(&x) // se envia la referencia de la memoria para que se modifique el valor no la copia
	// fmt.Println(x)

	p := Persona{"Piero", 35, "jabarca@soaint.com"}
	// acceder desde la estructura mediante la instancia
	p.diHola()

}

func editar(x *int) {
	*x = 20
}
