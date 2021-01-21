package projects

import (
	database "github.com/fdistorted/task_managment/db"
)

func DeleteById(userId, projectId string) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.Exec(database.DeleteProject, userId, projectId)

	if err != nil {
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
