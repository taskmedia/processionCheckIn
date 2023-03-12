package db

import (
	_ "github.com/lib/pq"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

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
