package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql" //la barra diagonal es de forma indirecta
)

func main() {
	// dns := "piero:admin@tcp(localhost:3306)/db_contacts" // por defecto en ese puerto se instala mysql

	// // Abrir una conexion a una base de datos
	// db, err := sql.Open("mysql", dns)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := db.Ping(); err != nil {
	// 	log.Fatal(err)
	// }

	// log.Println("Conexion a la base de datos MySql exitosa")

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Listar contactos:")
		fmt.Println("2. Obtener un usuario por ID")
		fmt.Println("3. Crear un usuario nuevo")
		fmt.Println("4. actualizar un usuario")
		fmt.Println("5. eliminar un usuario")

		// Leer la opcion seleccionada por el usuario

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			handlers.ListContacts(db)
		case 2:
			fmt.Println("Ingrese el ID del contacto")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.GetContactByID(db, idContact)
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
			handlers.ListContacts(db)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
			handlers.ListContacts(db)
		case 5:
			fmt.Println("Ingrese el ID del contacto que quiere eliminar")
			var idContact int
			fmt.Scanln(&idContact)
			handlers.DeleteContact(db, idContact)
			handlers.ListContacts(db)
		case 6:
			fmt.Println("Saliendo el programa ...")
			return
		default:
			fmt.Print("opcion incorrecta")

		}

	}

	// Crear un unstancia de Contact
	// newContact := models.Contact{
	// 	Name:  "Alex Bueno",
	// 	Email: "alex@soaint",
	// 	Phone: "945873291",
	// }

	// handlers.CreateContact(db, newContact)

	// modContact := models.Contact{
	// 	Id:    4,
	// 	Name:  "Jose Bueno",
	// 	Email: "jose@soaint",
	// 	Phone: "945873432",
	// }

	// handlers.UpdateContact(db, modContact)

	//handlers.DeleteContact(db, 5)

	// handlers.ListContacts(db)
	// handlers.GetContactByID(db, 2)

}

func inputContactDetails(option int) models.Contact {
	// Leer la entrada del usuario utilizando bufio
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact

	if option == 4 {
		fmt.Print("Ingrese ID del contacto: ")
		var idContact int
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}

	fmt.Print("Ingrese el nombre del contacto: ")
	name, _ := reader.ReadString('\n')
	contact.Name = strings.TrimSpace(name)

	fmt.Print("Ingrese el correo electronico del contacto: ")
	email, _ := reader.ReadString('\n')
	contact.Email = strings.TrimSpace(email)

	fmt.Print("Ingrese el numero de telefono del contacto: ")
	phone, _ := reader.ReadString('\n')
	contact.Phone = strings.TrimSpace(phone)

	return contact
}
