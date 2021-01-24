package tasks

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
	"time"
)

func GetAll(ctx context.Context, columnId, projectId string) (columns []models.Task, err error) {
	columns = make([]models.Task, 0)
	db := database.GetConn()
	defer db.Close()

	rows, err := db.QueryContext(ctx, database.GetAllTasks, projectId, columnId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id, title, description string
		var priority int
		var createdAt time.Time

		err = rows.Scan(&id, &title, &description, &priority, &createdAt)

		if err != nil {
			return
		}

		columns = append(columns, models.Task{
			Id:          id,
			Title:       title,
			Description: description,
			Priority:    &priority,
			Created:     createdAt,
		})

	}

	return
}
