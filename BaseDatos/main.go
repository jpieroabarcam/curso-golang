package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql" //la barra diagonal es de forma indirecta
)

func main() {
	dns := "piero:admin@tcp(localhost:3306)/db_contacts" // por defecto en ese puerto se instala mysql

	// Abrir una conexion a una base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Conexion a la base de datos MySql exitosa")
}
