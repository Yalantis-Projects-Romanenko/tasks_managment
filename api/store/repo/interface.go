package repo

import (
	"github.com/fdistorted/task_managment/store/repo/postgres/projects"
)

type Repo interface {
	Projects() projects.DAO
}
