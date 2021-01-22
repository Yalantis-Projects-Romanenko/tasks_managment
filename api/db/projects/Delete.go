package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
)

// TODO implement removing of all the rest
func DeleteById(userId, projectId string, ctx context.Context) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.ExecContext(ctx, database.DeleteProject, userId, projectId)

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
