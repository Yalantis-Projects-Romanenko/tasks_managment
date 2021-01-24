package tasks

import (
	"context"
	"errors"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func Update(ctx context.Context, task models.Task) (err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer database.RollbackWithHandler(ctx, tx)

	// update name if exists
	if len(task.Title) > 0 {
		// update name of task
		_, err = tx.ExecContext(ctx, database.UpdateTaskTitle, task.Title, task.Id, task.ProjectId)
		if err != nil {
			return
		}
	}

	// update name if exists
	if len(task.Description) > 0 {
		// update name of task
		_, err = tx.ExecContext(ctx, database.UpdateTaskDescription, task.Description, task.Id, task.ProjectId)
		if err != nil {
			return
		}
	}

	// update index if exists
	if task.Priority != nil {
		var maxIndex int
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetMaxTaskPriority, task.ColumnId, task.ProjectId).Scan(&maxIndex)
		if err != nil {
			return
		}

		if *task.Priority > maxIndex {
			err = database.ErrInvalidParameters.Wrap(errors.New("index cant be more then max index"))
			return
		}

		//getColumnIndex
		var currentPriority int
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetTasksPriorityById, task.ColumnId, task.ProjectId, task.Id).Scan(&currentPriority)
		if err != nil {
			return
		}

		if currentPriority == *task.Priority {
			err = database.ErrInvalidParameters.Wrap(errors.New("same index"))
			return
		}

		if currentPriority > *task.Priority {
			_, err = tx.ExecContext(ctx, database.IncrementTasksPriorities, task.ColumnId, task.ProjectId, task.Priority, currentPriority)
			if err != nil {
				return
			}
		} else {
			_, err = tx.ExecContext(ctx, database.DecrementTasksPriorities, task.ColumnId, task.ProjectId, task.Priority, currentPriority)
			if err != nil {
				return
			}
		}

		_, err = tx.ExecContext(ctx, database.UpdateTasksPriority, task.ColumnId, task.ProjectId, task.Id, task.Priority)
		if err != nil {
			return
		}
	}

	err = tx.Commit()
	return
}
