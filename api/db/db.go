package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

// Database variables
var connString string

func NewDb(connStr string) {
	connString = connStr
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
