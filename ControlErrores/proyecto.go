//gestor de contactos

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// Estructura de contactos
type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Guarda contactos en un archivo json
func saveContactsToFile(contacts []Contact) error {
	file, err := os.Create("contacts.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file) // codificar el slice a json (serializacion)
	err = encoder.Encode(contacts)
	if err != nil {
		return err
	}

	return nil
}

// Cargar contactos desde el archivo json
func loadContactsFromFile(contacts *[]Contact) error {
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	// Slice de contactos
	var contacts []Contact

	// Cargar contactos existentes desde  json

	err := loadContactsFromFile(&contacts)
	if err != nil {
		fmt.Println("Error al cargar el contacto")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("==== GESTOR DE CONTACTOS === \n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n")

		var option int
		_, err = fmt.Scanln(&option)
		if err != nil {
			fmt.Println("Error al leer la opcion:", err)
			return
		}

		switch option {
		case 1:
			// Ingresar y crfear contacto
			var c Contact
			fmt.Print("Nombre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			// Agregar un contacto a Slice
			contacts = append(contacts, c)

			// Guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto: ", err)
			}
		case 2:
			// Mostrar todos los contactos
			fmt.Println(("======================"))
			for index, contact := range contacts {
				fmt.Printf("%d Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println(("======================"))

		case 3:
			// Salir del programa
			return
		default:
			fmt.Println(" Opcion invalida")
		}

	}

}
