package tasks

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func GetById(ctx context.Context, projectId, columnId, taskId string) (*models.Task, error) {
	db := database.GetConn()
	defer db.Close()

	// create a user of models.User type
	var task models.Task

	row := db.QueryRowContext(ctx, database.GetTaskById, projectId, columnId, taskId)

	// unmarshal the row object to user
	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.Priority, &task.Created)
	if err != nil {
		return nil, err
	}

	return &task, nil
}
