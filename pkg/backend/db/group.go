package db

import (
	_ "github.com/lib/pq"

	// _ "github.com/taskmedia/processionCheckIn/pkg/backend/db/connection"
	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

func GetGroups() ([]model.Group, error) {
	query := "SELECT id, name FROM public.\"group\";"

	rows, err := DbConn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	groups := make([]model.Group, 0)

	for rows.Next() {
		var group model.Group
		if err := rows.Scan(&group.ID, &group.Name); err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
