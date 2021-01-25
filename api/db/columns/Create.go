package columns

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func CreateColumn(ctx context.Context, column models.Column) (columnId string, err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer database.RollbackWithHandler(ctx, tx)

	var index int
	// get list max index
	err = tx.QueryRowContext(ctx, database.GetMaxColumnIndex, column.ProjectId).Scan(&index)
	if err != nil {
		return
	}

	// create column
	err = tx.QueryRowContext(ctx, database.InsertColumn, column.Name, index+1, column.ProjectId).Scan(&columnId) // TODO set default column name via config
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}
