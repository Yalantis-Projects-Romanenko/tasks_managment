package tasks

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
)

func DeleteById(ctx context.Context, projectId, columnId, taskId string) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer database.RollbackWithHandler(ctx, tx)

	var currentPriority int
	// execute the sql statement
	err = tx.QueryRowContext(ctx, database.DeleteTaskById, projectId, columnId, taskId).Scan(&currentPriority)
	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, database.DecrementTasksPrioritiesOnDelete, projectId, columnId, currentPriority)
	if err != nil {
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return rowsAffected, nil
}
