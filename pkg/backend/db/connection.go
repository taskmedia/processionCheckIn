package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/taskmedia/processionCheckIn/pkg/cmd/parameter"
)

var DbConn *sql.DB

func init() {
	var err error
	sqlConnStr := fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		*parameter.PsqlDb,
		*parameter.PsqlUser,
		*parameter.PsqlPassword,
		*parameter.PsqlHost,
		*parameter.PsqlPort,
	)

	DbConn, err = sql.Open("postgres", sqlConnStr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: handle db close
	// defer DbConn.Close()
}
