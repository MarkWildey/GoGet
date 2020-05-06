package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var database *sql.DB

func main() {

	fmt.Print("Hello test1")

	database, _ = sql.Open("sqlite3", "a")

	statement, err := database.Prepare("create table a ( name int)") // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements

	database.Close()

	test2()

}
