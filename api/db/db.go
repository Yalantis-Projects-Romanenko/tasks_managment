package db

import (
	"database/sql"
	"fmt"
	"github.com/fdistorted/task_managment/config"
	"github.com/fdistorted/task_managment/logger"
	_ "github.com/lib/pq"
	"log"
)

// Database variables
var connString string

func NewDb(postgres config.Postgres) {
	//user=test password=test dbname=test sslmode=disable
	connString = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", postgres.Host, postgres.Port, postgres.Database, postgres.User, postgres.Password)

	logger.Get().Info(connString)
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
