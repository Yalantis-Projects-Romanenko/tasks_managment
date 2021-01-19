package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

//TODO use context in database queries
func Insert(userId string, project models.Project) (string, error) {
	db := database.GetConn()
	defer db.Close()

	var id string
	err := db.QueryRow("insert into projects (pname, pdescription, user_id) values($1,$2,$3) RETURNING id", project.Name, project.Description, userId).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
