package db

import (
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func CreateSeason(c *gin.Context) (int, error) {
	var season model.Season
	if err := c.ShouldBindJSON(&season); err != nil {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
		})

		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Error binding JSON in CreateSeason")

		return -1, err
	}

	if season.Name == "" {
		c.IndentedJSON(400, gin.H{
			"message": "Bad request",
			"error":   "name is empty",
		})

		log.Error("Error binding JSON in CreateSeason (name is empty)")

		return -1, errors.New("season name is empty")
	}

	err := DbConn.QueryRow(INSERT_SEASON, season.Name).Scan(&season.Id)
	if err != nil {
		return -1, err
	}

	return season.Id, nil
}

func GetSeasons() (interface{}, error) {
	rows, err := DbConn.Query(SELECT_SEASON_ALL)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	seasons := make([]model.Season, 0)

	for rows.Next() {
		var season model.Season
		if err := rows.Scan(&season.Id, &season.Name); err != nil {
			return nil, err
		}
		seasons = append(seasons, season)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return seasons, nil
}
