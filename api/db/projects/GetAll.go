package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
	"time"
)

const GetAllUsersProjects = "select id, pname, pdescription, created_at from projects where user_id = $1"

func GetAllByUserId(userId string) (projects []models.Project, err error) {
	db := database.GetConn()
	defer db.Close()

	rows, err := db.Query(GetAllUsersProjects, userId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id, pname, pdescription string
		var created_at time.Time

		err = rows.Scan(&id, &pname, &pdescription, &created_at)

		if err != nil {
			return
		}

		projects = append(projects, models.Project{
			Id:          id,
			Name:        pname,
			Description: pdescription,
			Created:     created_at,
		})

	}

	return projects, err
}
