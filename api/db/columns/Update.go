package columns

import (
	"context"
	"errors"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func Update(ctx context.Context, userId, projectId string, column models.Column) (err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// update name if exists
	if len(column.Name) > 0 {
		// update name of column
		_, err = tx.ExecContext(ctx, database.UpdateColumnName, column.Name, column.Id)
		if err != nil {
			// rollback if error
			tx.Rollback()
			return
		}
	}

	// update index if exists
	if column.Index != nil {
		var maxIndex int64
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetMaxColumnIndex, userId, projectId).Scan(&maxIndex)
		if err != nil {
			// rollback if error
			tx.Rollback()
			return
		}

		if *column.Index > maxIndex {
			err = database.ErrInvalidParameters.Wrap(errors.New("index cant be more then max index"))
			tx.Rollback()
			return
		}

		//getColumnIndex
		var currentIndex int64
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetColumnIndexById, userId, projectId, column.Id).Scan(&currentIndex)
		if err != nil {
			// rollback if error
			tx.Rollback()
			return
		}

		if currentIndex == *column.Index {
			err = database.ErrInvalidParameters.Wrap(errors.New("same index"))
			tx.Rollback()
			return
		}

		if currentIndex > *column.Index {
			_, err = tx.ExecContext(ctx, database.IncrementColumnIndexes, userId, projectId, column.Index, currentIndex)
			if err != nil {
				// rollback if error
				tx.Rollback()
				return
			}
		} else {
			_, err = tx.ExecContext(ctx, database.DecrementColumnIndexes, userId, projectId, column.Index, currentIndex)
			if err != nil {
				// rollback if error
				tx.Rollback()
				return
			}
		}

		_, err = tx.ExecContext(ctx, database.UpdateColumnIndex, userId, projectId, column.Id, column.Index)
		if err != nil {
			// rollback if error
			tx.Rollback()
			return
		}
	}

	err = tx.Commit()
	return
}
