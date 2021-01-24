package tasks

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func CreateTask(ctx context.Context, columnId, projectId string, column models.Task) (taskId string, err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer database.RollbackWithHandler(ctx, tx)

	var priority int
	// get list max priority
	err = tx.QueryRowContext(ctx, database.GetMaxTaskPriority, columnId, projectId).Scan(&priority)
	if err != nil {
		return
	}

	// create column
	err = tx.QueryRowContext(ctx, database.InsertTask, column.Title, column.Description, priority+1, projectId, columnId).Scan(&taskId) // TODO set default column name via config
	if err != nil {
		return
	}

	err = tx.Commit()
	return
}
