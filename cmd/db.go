package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	var connectionString string
	var dbType string

	dbType = "postgres"
	connectionString = "user=your_user password=your_password dbname=your_db host=your_host port=5432 sslmode=disable"

	db, err := sql.Open(dbType, connectionString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Успешное подключение к базе данных!")
}
