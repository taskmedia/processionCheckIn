package db

import (
	_ "github.com/lib/pq"

	"github.com/taskmedia/processionCheckIn/pkg/backend/db/model"
)

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
