package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"go.uber.org/zap"
)

func GetById(userId, projectId string) models.Project {
	db := database.GetConn()
	defer db.Close()

	// create a user of models.User type
	var project models.Project

	row := db.QueryRow("select id, pname, pdescription, created_at from projects where user_id = $1 and id = $2", userId, projectId)

	// unmarshal the row object to user
	err := row.Scan(&project.Id, &project.Name, &project.Description, &project.Created)
	if err != nil {
		logger.Get().Error("failed to scan sql result", zap.Error(err))
	}

	return project
}
