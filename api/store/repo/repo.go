package repo

import (
	"sync"

	"github.com/fdistorted/task_managment/store/repo/postgres/projects"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	projectsDAO projects.DAO
}

func Get() postgresRepo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			projectsDAO: projects.NewProjectsDAO(),
		}
	})

	return err
}

func (r postgresRepo) Projects() projects.DAO {
	return r.projectsDAO
}
