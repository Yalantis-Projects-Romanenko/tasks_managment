package columns

import (
	"context"
	database "github.com/fdistorted/task_managment/db"
	"github.com/fdistorted/task_managment/models"
)

func GetById(ctx context.Context, projectId, columnId string) (*models.Column, error) {
	db := database.GetConn()
	defer db.Close()

	// create a user of models.User type
	var column models.Column

	row := db.QueryRowContext(ctx, database.GetColumnById, projectId, columnId)

	// unmarshal the row object to user
	err := row.Scan(&column.Id, &column.Name, &column.Index, &column.Created)
	if err != nil {
		return nil, err
	}

	return &column, nil
}
