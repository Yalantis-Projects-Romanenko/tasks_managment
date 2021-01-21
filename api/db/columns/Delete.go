package columns

import (
	database "github.com/fdistorted/task_managment/db"
)

func DeleteById(userId, projectId, columnId string) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.Exec(database.DeleteColumnById, userId, projectId, columnId)
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
