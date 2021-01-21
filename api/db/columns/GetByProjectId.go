package columns

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
	"time"
)

func GetAll(userId, projectId string) (columns []models.Column, err error) {
	columns = make([]models.Column, 0)
	db := database.GetConn()
	defer db.Close()

	rows, err := db.Query(database.GetAllUsersColumnsByProject, userId, projectId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var index int64
		var createdAt time.Time

		err = rows.Scan(&id, &name, &index, &createdAt)

		if err != nil {
			return
		}

		columns = append(columns, models.Column{
			Id:      id,
			Name:    name,
			Index:   index,
			Created: createdAt,
		})

	}

	return
}
