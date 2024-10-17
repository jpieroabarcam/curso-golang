package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// ListContacts  lista todos los contactos desde la base de datos
func ListContacts(db *sql.DB) {
	//Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// iterar sore los resultados y mostrarlos
	fmt.Println("\n LISTA DE CONTACTOS:")
	fmt.Println("----------------------------------------------------")
	for rows.Next() {
		//Instancia de modelo contact
		contact := models.Contact{}
		var valueEmail sql.NullString
		err := rows.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
		if err != nil {
			log.Fatal(err)
		}

		if valueEmail.Valid {
			contact.Email = valueEmail.String
		} else {
			contact.Email = "Sin correo electronico"
		}
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)

		fmt.Println("----------------------------------------------------")

	}

}

// GetContactByID obtiene un contacto de la base de datos mediante su ID
func GetContactByID(db *sql.DB, contactID int) {
	query := "SELECT * FROM contact WHERE id = ?"

	row := db.QueryRow(query, contactID)

	// Instancia de modelo contact
	contact := models.Contact{}
	var valueEmail sql.NullString // para manejar el valor null

	// escanear el valor de contact

	err := row.Scan(&contact.Id, &contact.Name, &valueEmail, &contact.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("No se encontro ningun contacto con el ID %d", contactID)
		}
	}

	if valueEmail.Valid {
		contact.Email = valueEmail.String
	} else {
		contact.Email = "Sin correo electronico"
	}

	fmt.Println("\n LISTA DE UN CONTACTO:")
	fmt.Println("----------------------------------------------------")

	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)

	fmt.Println("----------------------------------------------------")
}

// CreateContact registra un nuevo contacto en la base de datos
func CreateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para insertar un nuevo contacto
	query := "INSERT INTO contact(name, email, phone) VALUES(?, ?, ?)"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Nuevo contacto registrado con exito")
}

// UpdateContact actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact) {
	// Sentencia SQL para actualizar u contacto
	query := "UPDATE contact SET name = ?, email = ?, phone=? WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto actualizado con exito")

}

// DeleteContact elimina un contacto existente en la base de datos
func DeleteContact(db *sql.DB, contactId int) {
	// Sentencia SQL para eliminar un contacto
	query := "DELETE FROM contact WHERE id = ?"

	// Ejecutar la sentencia SQL
	_, err := db.Exec(query, contactId)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Contacto ELIMINADO con exito")
}
