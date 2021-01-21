package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

const InsertProject = "insert into projects (pname, pdescription, user_id) values($1,$2,$3) RETURNING id"

//TODO use context in database queries
func Insert(project models.Project) (string, error) {
	db := database.GetConn()
	defer db.Close()

	var id string
	err := db.QueryRow(InsertProject, project.Name, project.Description, project.UserId).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
