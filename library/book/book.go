package book

import "fmt"

//Polimorfismo
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

//usar con mayuscula es en mayuscula, caso contrario en minuscula para ser privado
type Book struct {
	title  string
	author string
	pages  int
}

//Simular un constructor - creando una funcion
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// metodos set y get para los atributos privados
func (b *Book) SetTitle(title string) {
	b.title = title
}

func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) PrintInfo() {
	fmt.Printf(" Title: %s\n Author: %s\n Pages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

func NewTextBook(title, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages}, //Composicion
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf(" Title: %s\n Author: %s\n Pages: %d\n Editorial: %s\n Nivel: %s\n",
		b.title, b.author, b.pages, b.editorial, b.level)
}
