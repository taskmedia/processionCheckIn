package db

func dropDb() error {
	err := executePsqlFile("./psql/drop.psql")
	if err != nil {
		return err
	}

	return nil
}

func InitDb() error {
	err := executePsqlFile("./psql/init-table.psql")
	if err != nil {
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
