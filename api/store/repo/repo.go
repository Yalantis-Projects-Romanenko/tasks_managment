package repo

import (
	"sync"

	"github.com/fdistorted/task_managment/store/repo/postgres/projects"
	"github.com/fdistorted/task_managment/store/repo/postgres/symbols"
)

var (
	repo postgresRepo
	once = &sync.Once{}
)

type postgresRepo struct {
	symbolsDAO symbols.DAO
	candlesDAO projects.DAO
}

func Get() Repo {
	return repo
}

func Load() (err error) {
	once.Do(func() {
		repo = postgresRepo{
			symbolsDAO: symbols.NewSymbolsDAO(),
			candlesDAO: projects.NewCandlesDAO(),
		}
	})

	return err
}

func (r postgresRepo) Symbols() symbols.DAO {
	return r.symbolsDAO
}

func (r postgresRepo) Candles() projects.DAO {
	return r.candlesDAO
}
