package db

import (
	_ "github.com/lib/pq"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func GetLocations() (interface{}, error) {
	query := "SELECT id, name FROM public.\"location\";"

	rows, err := DbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	locations := make([]model.Location, 0)

	for rows.Next() {
		var location model.Location
		if err := rows.Scan(&location.ID, &location.Name); err != nil {
			return nil, err
		}
		locations = append(locations, location)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return locations, nil
}