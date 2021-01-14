package projects

import (
	"context"
	"github.com/fdistorted/task_managment/models"
	"github.com/fdistorted/task_managment/store/repo/postgres"
)

type projectsDAO struct {
	q *postgres.DBQuery
}

func NewProjectsDAO() DAO {
	return &projectsDAO{q: postgres.GetDB().QueryContext(context.Background())}
}

func (dao projectsDAO) WithTx(tx *postgres.DBQuery) DAO {
	return &projectsDAO{q: tx}
}

func (dao projectsDAO) Get(id int64) (project models.Project, err error) {
	err = dao.q.Model(&project).
		Where("id = ?", id).
		Select()

	return project, err
}

func (dao projectsDAO) Insert(project models.Project) (_ models.Project, err error) {
	_, err = dao.q.Model(&project).
		Returning("*").
		Insert()

	return project, err
}
