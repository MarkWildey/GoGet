package main

import (
	"database/sql"
	"fmt"
	_ "fmt"
	"log"
)

func test2() {

	fmt.Print("Hello test1")

	database, _ = sql.Open("sqlite3", "b")

	statement, err := database.Prepare("create table b ( name int)") // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements

	database.Close()

}
