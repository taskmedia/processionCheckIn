package db

import (
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

func executePsqlFile(filename string) {
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
	_, err = DbConn.Exec(string(commands))
	if err != nil {
		panic(err)
	}

	fmt.Println("SQL commands executed successfully.")
}
