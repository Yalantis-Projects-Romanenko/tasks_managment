package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"go.uber.org/zap"
)

//TODO use context in database queries
func CreateProject(project models.Project, ctx context.Context) (id string, err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	// create project
	err = tx.QueryRowContext(ctx, database.InsertProject, project.Name, project.Description, project.UserId).Scan(&id)
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	var columnId string
	// create column
	err = tx.QueryRowContext(ctx, database.InsertColumn, "Default", 0, id).Scan(&columnId) // TODO set default column name via config
	if err != nil {
		// rollback if error
		tx.Rollback()
		return
	}

	logger.WithCtxValue(ctx).Debug("created column", zap.String("columnId", columnId))

	err = tx.Commit()
	return
}
