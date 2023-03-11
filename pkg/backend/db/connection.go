package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DbConn *sql.DB

func init() {
	var err error
	DbConn, err = sql.Open("postgres", "dbname=postgres user=postgres password=7z2Czy8s61 host=localhost port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: handle db close
	// defer DbConn.Close()
}
