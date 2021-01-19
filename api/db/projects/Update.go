package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"go.uber.org/zap"
	"log"
)

func Update(userId string, project models.Project) int64 {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.Exec(`UPDATE projects SET pname=$3, pdescription=$4 WHERE user_id=$1 and id=$2`, userId, project.Id, project.Name, project.Description)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	logger.Get().Info("Total rows/record affected %v", zap.Int64("rowsAffected", rowsAffected))

	return rowsAffected
}
