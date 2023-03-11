package db

func dropDb() {
	executePsqlFile("./psql/drop.psql")
}

func InitDb() {
	executePsqlFile("./psql/init-table.psql")
}

func ResetDb() {
	dropDb()
	InitDb()
}

func ResetDbExampledata() {
	dropDb()
	InitDb()

	executePsqlFile("./psql/exampledata.psql")
}
