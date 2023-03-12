package db

import (
	_ "embed"
)

// Embed the SQL scripts into the binary via variable

//go:embed scripts/drop.psql
var PSQL_SCRIPT_DROP string

//go:embed scripts/exampledata.psql
var PSQL_SCRIPT_EXAMPLEDATA string

//go:embed scripts/init-table.psql
var PSQL_SCRIPT_INITTABLE string

func dropDb() error {
	_, err := DbConn.Exec(PSQL_SCRIPT_DROP)
	return err
}

func InitDb() error {
	_, err := DbConn.Exec(PSQL_SCRIPT_INITTABLE)
	return err
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

	_, err = DbConn.Exec(PSQL_SCRIPT_EXAMPLEDATA)
	return err
}
