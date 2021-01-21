package columns

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

//TODO use context in database queries
func CreateColumn(userId, projectId string, column models.Column, ctx context.Context) (columnId string, err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	var index int
	// get list max index
	err = tx.QueryRowContext(ctx, database.GetMaxColumnIndex, userId, projectId).Scan(&index)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	// create column
	err = tx.QueryRowContext(ctx, database.InsertColumn, column.Name, index+1, projectId).Scan(&columnId) // TODO set default column name via config
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	err = tx.Commit()
	return
}
