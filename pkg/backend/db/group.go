package db

import (
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func CreateGroup(c *gin.Context) (int, error) {
	var group model.Group
	if err := c.ShouldBindJSON(&group); err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error binding JSON in CreateGroup")

		return -1, err
	}

	if group.Name == "" {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
			"error":   "name is empty",
		})

		log.Error("Error binding JSON in CreateGroup (name is empty)")

		return -1, errors.New("group name is empty")
	}

	err := DbConn.QueryRow(INSERT_GROUP, group.Name).Scan(&group.Id)
	if err != nil {
		return -1, err
	}

	return group.Id, nil
}

func GetGroups() (interface{}, error) {
	rows, err := DbConn.Query(SELECT_GROUP_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := make([]model.Group, 0)

	for rows.Next() {
		var group model.Group
		if err := rows.Scan(&group.Id, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
