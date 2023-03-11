package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func GetUsers() ([]model.User, error) {
	query := "SELECT id, firstname, lastname FROM public.\"user\";"

	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=7z2Czy8s61 host=localhost port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]model.User, 0)

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUser(id int) (model.User, error) {
	query := "SELECT id, firstname, lastname FROM public.\"user\" WHERE id = $1;"

	db, err := sql.Open("postgres", "dbname=postgres user=postgres password=7z2Czy8s61 host=localhost port=5432 sslmode=disable")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer db.Close()

	var user model.User

	if err := db.QueryRow(query, id).Scan(&user.ID, &user.Firstname, &user.Lastname); err != nil {
		return user, err
	}

	return user, nil
}
