package db

import (
	log "github.com/sirupsen/logrus"
)

func dropDb() error {
	err := executePsqlFile("./psql/drop.psql")
	if err != nil {

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error dropping database")

		return err
	}

	return nil
}

func InitDb() error {
	err := executePsqlFile("./psql/init-table.psql")
	if err != nil {

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error initializing database")

		return err
	}

	return nil
}

func ResetDb() error {
	err := dropDb()
	if err != nil {
		return err
	}

	return InitDb()
}

func ResetDbExampledata() error {
	err := ResetDb()
	if err != nil {
		return err
	}

	return executePsqlFile("./psql/exampledata.psql")
}
