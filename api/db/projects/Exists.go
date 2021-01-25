package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
)

func Exists(ctx context.Context, userId, projectId string) (bool, error) {
	db := database.GetConn()
	defer db.Close()

	exists := false

	err := db.QueryRowContext(ctx, database.ExistsProject, userId, projectId).Scan(&exists)

	if err != nil {
		return exists, err
	}

	return exists, nil
}
