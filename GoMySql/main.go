package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()

	fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	//db.TrunctateTable("users")
	user := models.CreateUser("piero", "44231", "piero@soaint.com")
	fmt.Println(user)
	// users := models.ListUsers()
	// fmt.Println(users)
	// user := models.GetUser(2)
	// fmt.Println(user)

	// user.Delete()
	//db.TrunctateTable("users")
	users := models.ListUsers()
	fmt.Println(users)

	db.Close()
	//db.Ping()
}
