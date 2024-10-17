package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	// Cadena de conexion a la base de datos MySQL

	dns := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Abrir una conexion a una base de datos
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
		//log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
		//log.Fatal(err)
	}

	log.Println("Conexion a la base de datos MySql exitosa")

	return db, nil
}
