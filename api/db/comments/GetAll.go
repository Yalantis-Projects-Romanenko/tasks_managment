package comments

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
	"time"
)

func GetAll(ctx context.Context, projectId, taskId string) (comments []models.Comment, err error) {
	comments = make([]models.Comment, 0)
	db := database.GetConn()
	defer db.Close()

	rows, err := db.QueryContext(ctx, database.GetAllCommentsOfTheTask, projectId, taskId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id, username, content string
		var createdAt time.Time

		err = rows.Scan(&id, &username, &content, &createdAt)

		if err != nil {
			return
		}

		comments = append(comments, models.Comment{
			Id:      id,
			UserId:  username,
			Content: content,
			Created: createdAt,
		})

	}

	return
}
