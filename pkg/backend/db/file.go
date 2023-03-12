package db

import (
	"os"

	_ "github.com/lib/pq"
)

func executePsqlFile(filename string) error {
	commands, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	_, err = DbConn.Exec(string(commands))
	if err != nil {
		return err
	}

	return nil
}
