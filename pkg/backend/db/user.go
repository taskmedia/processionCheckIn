package db

import (
	_ "github.com/lib/pq"
	// _ "github.com/taskmedia/processionCheckIn/pkg/backend/db/connection"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func GetUser(id int) (model.User, error) {
	query := "SELECT id, firstname, lastname FROM public.\"user\" WHERE id = $1;"

	var user model.User

	if err := DbConn.QueryRow(query, id).Scan(&user.ID, &user.Firstname, &user.Lastname); err != nil {
		return user, err
	}

	return user, nil
}

func GetUsers() ([]model.User, error) {
	query := "SELECT id, firstname, lastname FROM public.\"user\";"

	rows, err := DbConn.Query(query)
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
