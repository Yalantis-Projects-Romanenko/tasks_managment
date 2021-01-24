package tasks

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
)

func DeleteById(ctx context.Context, projectId, columnId, taskId string) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.ExecContext(ctx, database.DeleteTaskById, projectId, columnId, taskId)
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
