package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/logger"
	"github.com/fdistorted/task_managment/models"
	"go.uber.org/zap"
)

//todo use context in databse queries
func Insert(userId, project models.Project) error {
	db := database.GetConn()
	tx, err := db.Begin()
	if err != nil {
		logger.Get().Fatal("failed to start transaction", zap.Error(err))
	}

	defer tx.Rollback()

	return nil

	//stmnt, err = tx.ExecContext("insert into projects (projects pname, pdescription, user_id) values($1,$2,$3)")
	//
	//rows, err := db.Query("select id, pname, pdescription, created_at from project where user_id = $1", userId)
	//
	//if err != nil {
	//	logger.Get().Fatal("Cannot connect: ", zap.Error(err))
	//}
	//
	//defer rows.Close()
	//
	//for rows.Next() {
	//	var id, pname, pdescription string
	//	var created_at time.Time
	//
	//	err = rows.Scan(&id, &pname, &pdescription, &created_at)
	//	if err != nil {
	//		log.Printf(err.Error())
	//	}
	//
	//	projects = append(projects, models.Project{
	//		Id:          id,
	//		Name:        pname,
	//		Description: pdescription,
	//		Created:     created_at,
	//	})
	//
	//}
	//
	//return projects
}
