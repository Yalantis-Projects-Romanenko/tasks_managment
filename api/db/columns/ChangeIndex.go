package columns

import (
	"context"
	"errors"
	database "github.com/fdistorted/task_managment/db"
)

// ChangeIndex moves column inside of a project
func ChangeIndex(ctx context.Context, userId, projectId, columnId string, index int64) (err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	var maxIndex int64
	// get list max index
	err = tx.QueryRowContext(ctx, database.GetMaxColumnIndex, userId, projectId).Scan(&maxIndex)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	if index > maxIndex {
		err = database.ErrInvalidParameters.Wrap(errors.New("index cant be more then max index"))
		tx.Rollback()
		return
	}

	//getColumnIndex
	var currentIndex int64
	// get list max index
	err = tx.QueryRowContext(ctx, database.GetColumnIndexById, userId, projectId, columnId).Scan(&currentIndex)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	if currentIndex == index {
		err = database.ErrInvalidParameters.Wrap(errors.New("same index"))
		tx.Rollback()
		return
	}

	if currentIndex > index {
		_, err = tx.ExecContext(ctx, database.IncrementColumnIndexes, userId, projectId, index, currentIndex)
		if err != nil {
			// rollback if error
			tx.Rollback()
			return
		}
	} else {
		_, err = tx.ExecContext(ctx, database.DecrementColumnIndexes, userId, projectId, index, currentIndex)
		if err != nil {
			// rollback if error
			tx.Rollback()
			return
		}
	}

	_, err = tx.ExecContext(ctx, database.UpdateColumnIndex, userId, projectId, columnId, index)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
