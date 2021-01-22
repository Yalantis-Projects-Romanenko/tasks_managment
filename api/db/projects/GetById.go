package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func GetById(userId, projectId string, ctx context.Context) (*models.Project, error) {
	db := database.GetConn()
	defer db.Close()

	// create a user of models.User type
	var project models.Project

	row := db.QueryRowContext(ctx, database.GetProjectById, userId, projectId)

	// unmarshal the row object to user
	err := row.Scan(&project.Id, &project.Name, &project.Description, &project.Created)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
