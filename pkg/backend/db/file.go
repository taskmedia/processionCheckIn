package db

import (
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func executePsqlFile(filename string) error {
	// Read SQL file
	commands, err := os.ReadFile(filename)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error reading SQL file")

		return err
	}

	// Execute SQL file
	_, err = DbConn.Exec(string(commands))
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error executing SQL file")

		return err
	}

	return nil
}
