package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//username:password@tcp(localhost:3306)/database
const url = "root:root@tcp(localhost:3306)/goweb_db"

// Guarda la conexion
var db *sql.DB

// Realiza la conexion
func Connect() {
	conection, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	db = conection

}

// Cerrar la conexion
func Close() {
	db.Close()
}

// Verificar la conexion
func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

// Crea una tabla
func CreateTable(schema string, name string) {

	if !ExistTable(name) {
		_, err := db.Exec(schema)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}

// Verifica si existe una tabla o no
func ExistTable(tableName string) bool {
	sql := fmt.Sprintf("SHOW TABLES LIKE '%s' ", tableName)

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return rows.Next()
}

// Reiniciar el registro de una table
func TrunctateTable(tableName string) {
	sql := fmt.Sprintf("TRUNCATE %s", tableName)
	Exec(sql)
}

// Polimorfismo exec
func Exec(query string, args ...any) (sql.Result, error) {
	result, err := db.Exec(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

// Polimorfismo query
func Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println(err)
	}
	return rows, err
}
