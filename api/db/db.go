package db

import (
	"database/sql"
	"github.com/fdistorted/task_managment/config"
	"github.com/fdistorted/task_managment/logger"
	_ "github.com/lib/pq"
	"log"
)

// Database variables
var connString string

func NewDb(postgres config.Postgres) {
	connString = "postgresql://" + postgres.User + ":" + postgres.Password + "@" + postgres.Host + ":" + postgres.Port + "/" + postgres.Database
	logger.Get().Debug(connString)
}

func GetConn() *sql.DB {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Cannot connect: ", err.Error())
	}

	return db
}
