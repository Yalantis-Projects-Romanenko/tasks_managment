package columns

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
)

func DeleteById(ctx context.Context, projectId, columnId string) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}

	defer database.RollbackWithHandler(ctx, tx)

	var currentIndex int
	// execute the sql statement
	err = tx.QueryRowContext(ctx, database.DeleteColumnById, projectId, columnId).Scan(&currentIndex)
	if err != nil {
		return 0, err
	}

	res, err := tx.ExecContext(ctx, database.DecrementColumnIndexesOnDelete, projectId, currentIndex)
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
