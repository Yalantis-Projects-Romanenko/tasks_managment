package tasks

import (
	"context"
	"errors"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func Update(ctx context.Context, columnId string, task models.Task) (err error) {
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
		_, err = tx.ExecContext(ctx, database.UpdateTaskTitle, task.Title, task.Id, task.ProjectId, columnId)
		if err != nil {
			return
		}
	}

	// update name if exists
	if len(task.Description) > 0 {
		// update name of task
		_, err = tx.ExecContext(ctx, database.UpdateTaskDescription, task.Description, task.Id, task.ProjectId, columnId)
		if err != nil {
			return
		}
	}

	//change task status and priority
	if task.ColumnId != "" {
		priorityToSet := 0

		var maxPriority int
		// get max task priority
		err = tx.QueryRowContext(ctx, database.GetMaxTaskPriority, task.ColumnId, task.ProjectId).Scan(&maxPriority)
		if err != nil {
			return
		}

		if task.Priority != nil {
			if *task.Priority < 0 {
				return errors.New("wrong task priority")
			}

			if *task.Priority > maxPriority {
				err = database.ErrInvalidParameters.Wrap(errors.New("priority can't be more then max index"))
				return
			}

			if *task.Priority >= maxPriority+1 {
				priorityToSet = maxPriority + 1
			}

		}

		if priorityToSet < maxPriority+1 {
			_, err = tx.ExecContext(ctx, database.IncrementTasksPrioritiesOnMove, task.ProjectId, task.ColumnId, priorityToSet)
			if err != nil {
				return
			}
		}
		// update name of task
		_, err = tx.ExecContext(ctx, database.UpdateTaskColumnAndPriority, task.ColumnId, priorityToSet, task.Id, task.ProjectId)
		if err != nil {
			return
		}
	} else if task.Priority != nil { // update only priority
		var maxPriority int
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetMaxTaskPriority, columnId, task.ProjectId).Scan(&maxPriority)
		if err != nil {
			return
		}

		if *task.Priority > maxPriority {
			err = database.ErrInvalidParameters.Wrap(errors.New("index cant be more then max index"))
			return
		}

		//getColumnIndex
		var currentPriority int
		// get list max index
		err = tx.QueryRowContext(ctx, database.GetTasksPriorityById, columnId, task.ProjectId, task.Id).Scan(&currentPriority)
		if err != nil {
			return
		}

		if currentPriority == *task.Priority {
			err = database.ErrInvalidParameters.Wrap(errors.New("same index"))
			return
		}

		if currentPriority > *task.Priority {
			_, err = tx.ExecContext(ctx, database.IncrementTasksPriorities, columnId, task.ProjectId, task.Priority, currentPriority)
			if err != nil {
				return
			}
		} else {
			_, err = tx.ExecContext(ctx, database.DecrementTasksPriorities, columnId, task.ProjectId, task.Priority, currentPriority)
			if err != nil {
				return
			}
		}

		_, err = tx.ExecContext(ctx, database.UpdateTasksPriority, columnId, task.ProjectId, task.Id, task.Priority)
		if err != nil {
			return
		}
	}

	err = tx.Commit()
	return
}
