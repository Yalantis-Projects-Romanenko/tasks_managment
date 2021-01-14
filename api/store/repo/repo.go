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
	candlesDAO projects.DAO
}

func Get() postgresRepo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			candlesDAO: projects.NewCandlesDAO(),
		}
	})

	return err
}

func (r postgresRepo) Candles() projects.DAO {
	return r.candlesDAO
}
