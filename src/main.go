package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // this is how database/sql recognizes postgres driver
)

func main() {
	connStr := "user=al password=pw dbname=words sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// todo -- read from combined files and put into db
	rows, err := db.Query("INSERT INTO words (string) VALUES ($1)", "abajo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println(rows)
}
