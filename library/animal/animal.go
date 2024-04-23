package animal

import "fmt"

type Animal interface {
	Sonido()
}

// estructura perro y sus metodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido() {
	fmt.Println(p.Nombre + " hace guau guau")
}

// estructura gato y sus metodos
type Gato struct {
	Nombre string
}

func (p *Gato) Sonido() {
	fmt.Println(p.Nombre + " hace miau")
}

// funcion para imprimir sonido
func HacerSonido(animal Animal) {
	animal.Sonido()
}
