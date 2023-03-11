package db

import (
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func checkUserExists(id int) bool {
	var user model.User

	if err := DbConn.QueryRow(SELECT_USERID_BY_ID, id).Scan(&user.ID); err != nil {
		return false
	}

	return true
}

func CreateUser(user model.User) (model.User, error) {
	err := DbConn.QueryRow(INSERT_USER, user.Firstname, user.Lastname).Scan(&user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func DeleteUser(id int) error {
	// check if user exists
	if !checkUserExists(id) {
		return fmt.Errorf("id does not exist")
	}

	result, err := DbConn.Exec(DELETE_USER_BY_ID, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// check if user was deleted
	if rowsAffected == 0 {
		return fmt.Errorf("db.DeleteUser did not delete any rows")
	}

	return nil
}

func GetUser(id int) (interface{}, error) {
	var user model.User

	if err := DbConn.QueryRow(SELECT_USER_BY_ID, id).Scan(&user.ID, &user.Firstname, &user.Lastname); err != nil {
		log.WithFields(logrus.Fields{
			"error": err.Error(),
		}).Error("Error getting user from database in GetUser")

		return user, err
	}

	return user, nil
}

func GetUsers() (interface{}, error) {
	rows, err := DbConn.Query(SELECT_USER_ALL)
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
