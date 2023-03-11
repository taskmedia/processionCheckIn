package db

import (
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func executePsqlFile(filename string) error {
	// Open SQL file
	file, err := os.Open(filename)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error opening SQL file")

		return err
	}
	defer file.Close()

	// Read SQL file
	commands, err := ioutil.ReadAll(file)
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
