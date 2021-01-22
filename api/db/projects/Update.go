package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func Update(ctx context.Context, project models.Project) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.ExecContext(ctx, database.UpdateProject, project.UserId, project.Id, project.Name, project.Description)

	if err != nil {
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
