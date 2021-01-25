package comments

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func CreateComment(ctx context.Context, comment models.Comment) (columnId string, err error) {
	db := database.GetConn()
	defer db.Close()

	// create comment
	err = db.QueryRowContext(ctx, database.InsertComment, comment.ProjectId, comment.TaskId, comment.UserId, comment.Content).Scan(&columnId)
	if err != nil {
		return
	}
	return
}
