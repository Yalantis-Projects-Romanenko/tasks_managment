package projects

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
	"time"
)

func GetAllByUserId(ctx context.Context, userId string) (projects []models.Project, err error) {
	projects = make([]models.Project, 0) //TODO use this thing in all requests or check it with mentor
	db := database.GetConn()
	defer db.Close()

	rows, err := db.QueryContext(ctx, database.GetAllUsersProjects, userId)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id, pname, pdescription string
		var createdAt time.Time
		err = rows.Scan(&id, &pname, &pdescription, &createdAt)
		if err != nil {
			return
		}
		projects = append(projects, models.Project{
			Id:          id,
			Name:        pname,
			Description: pdescription,
			Created:     createdAt,
		})
	}

	return projects, err
}
