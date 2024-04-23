package main

import (
	"fmt"
	"library/animal"
	"library/book"
)

func main() {

	//a partir del "constructor"
	myBook := book.NewBook("Los Cachorros", "Mario Vargas Llosa", 900)

	// var myBook = book.Book{
	// 	Title:  "O Alquimista",
	// 	Author: "Paulo Coelho",
	// 	Pages:  300,
	// }
	//myBook.PrintInfo()

	myBook.SetTitle("La Ciudad y los Perros")
	//fmt.Println(myBook.GetTitle())
	//myBook.PrintInfo()

	myTextBook := book.NewTextBook("Fisica", "Serway", 1000, "Mc Graw Hill", "Superior")
	//myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myTextBook)

	fmt.Println("======================")

	// crear objetos animal
	miPerro := animal.Perro{Nombre: "Boby"}
	miGato := animal.Gato{Nombre: "Chizito"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Tyson"},
		&animal.Gato{Nombre: "Alex"},
		&animal.Perro{Nombre: "Lucas"},
		&animal.Gato{Nombre: "Luna"},
	}

	fmt.Println("======================")

	for _, animal := range animales {
		animal.Sonido()
	}

}
