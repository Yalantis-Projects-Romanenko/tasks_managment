package columns

import (
	"context"
	"errors"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func Update(ctx context.Context, column models.Column) (err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer database.RollbackWithHandler(ctx, tx)

	// update name if exists
	if len(column.Name) > 0 {
		// update name of column
		_, err = tx.ExecContext(ctx, database.UpdateColumnName, column.ProjectId, column.Id, column.Name)
		if err != nil {
			return
		}
	}

	// update index if exists
	if column.Index != nil {

		//getColumnIndex
		var currentIndex int64
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetColumnIndexById, column.ProjectId, column.Id).Scan(&currentIndex)
		if err != nil {
			return
		}

		if currentIndex == *column.Index {
			err = tx.Commit()
			return err
			//err = database.ErrInvalidParameters.Wrap(errors.New("same index"))
			//return
		}

		var maxIndex int64
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetMaxColumnIndex, column.ProjectId).Scan(&maxIndex)
		if err != nil {
			return
		}

		if *column.Index > maxIndex {
			err = database.ErrInvalidParameters.Wrap(errors.New("index cant be more then max index"))
			return
		}

		if currentIndex > *column.Index {
			_, err = tx.ExecContext(ctx, database.IncrementColumnIndexes, column.ProjectId, column.Index, currentIndex)
			if err != nil {
				return
			}
		} else {
			_, err = tx.ExecContext(ctx, database.DecrementColumnIndexes, column.ProjectId, column.Index, currentIndex)
			if err != nil {
				return
			}
		}

		_, err = tx.ExecContext(ctx, database.UpdateColumnIndex, column.ProjectId, column.Id, column.Index)
		if err != nil {
			return
		}
	}

	err = tx.Commit()
	return
}
