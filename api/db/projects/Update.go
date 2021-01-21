package projects

import (
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

const UpdateProject = `UPDATE projects SET pname=$3, pdescription=$4 WHERE user_id=$1 and id=$2`

func Update(project models.Project) (int64, error) {
	db := database.GetConn()
	defer db.Close()

	// execute the sql statement
	res, err := db.Exec(UpdateProject, project.UserId, project.Id, project.Name, project.Description)

	if err != nil {
		return 0, err
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
