package db

import (
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func CreateLocation(c *gin.Context) (int, error) {
	var location model.Location
	if err := c.ShouldBindJSON(&location); err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error binding JSON in CreateLocation")

		return -1, err
	}

	if location.Name == "" {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
			"error":   "name is empty",
		})

		log.Error("Error binding JSON in CreateLocation (name is empty)")

		return -1, errors.New("location name is empty")
	}

	err := DbConn.QueryRow(INSERT_LOCATION, location.Name).Scan(&location.Id)
	if err != nil {
		return -1, err
	}

	return location.Id, nil
}

func GetLocations() (interface{}, error) {
	rows, err := DbConn.Query(SELECT_LOCATION_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := make([]model.Location, 0)

	for rows.Next() {
		var location model.Location
		if err := rows.Scan(&location.Id, &location.Name); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}
