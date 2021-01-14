package projects

import (
	"github.com/fdistorted/task_managment/models"
)

type DAO interface {
	Get(id string) (models.Project, error)

	Insert(project models.Project) (_ models.Project, err error)
}
