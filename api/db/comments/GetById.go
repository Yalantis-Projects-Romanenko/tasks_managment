package comments

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func GetById(ctx context.Context, projectId, taskId, commentId string) (*models.Comment, error) {
	db := database.GetConn()
	defer db.Close()

	// create a user of models.User type
	var comment models.Comment

	row := db.QueryRowContext(ctx, database.GetCommentById, projectId, taskId, commentId)

	// unmarshal the row object to user
	err := row.Scan(&comment.Id, &comment.UserId, &comment.Content, &comment.Created)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}
