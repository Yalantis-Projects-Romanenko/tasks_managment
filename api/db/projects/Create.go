package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"go.uber.org/zap"
)

//TODO use context in database queries
func CreateProject(ctx context.Context, project models.Project) (id string, err error) {
	db := database.GetConn()
	defer db.Close()

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer database.RollbackWithHandler(ctx, tx)

	// create project
	err = tx.QueryRowContext(ctx, database.InsertProject, project.Name, project.Description, project.UserId).Scan(&id)
	if err != nil {
		return
	}

	var columnId string
	// create column
	err = tx.QueryRowContext(ctx, database.InsertColumn, "Default", 0, id).Scan(&columnId) // TODO set default column name via config
	if err != nil {
		return
	}

	logger.WithCtxValue(ctx).Debug("created column", zap.String("columnId", columnId))

	err = tx.Commit()
	return
}
