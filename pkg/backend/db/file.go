package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func executePsqlFile(filename string) {
	// Open a database connection
	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=7z2Czy8s61 host=localhost port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Read SQL commands from file
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Execute SQL commands
	commands, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(string(commands))
	if err != nil {
		panic(err)
	}

	fmt.Println("SQL commands executed successfully.")
}
