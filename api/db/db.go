package db

import (
	"database/sql"
	"fmt"
	"github.com/fdistorted/task_managment/config"
	"github.com/fdistorted/task_managment/logger"
	_ "github.com/lib/pq"
	"log"
)

const (
	DeleteProject       = `DELETE FROM projects WHERE user_id=$1 and id=$2`
	GetAllUsersProjects = "select id, pname, pdescription, created_at from projects where user_id = $1"
	GetProjectById      = "select id, pname, pdescription, created_at from projects where user_id = $1 and id = $2 limit 1"
	InsertProject       = "insert into projects (pname, pdescription, user_id) values($1,$2,$3) RETURNING id"
	UpdateProject       = `UPDATE projects SET pname=$3, pdescription=$4 WHERE user_id=$1 and id=$2`

	InsertColumn                = "insert into lists (cname, index, project_id) values($1,$2,$3) RETURNING id"
	GetMaxColumnIndex           = "select MAX (lists.index) from projects left join lists on lists.project_id = projects.id where projects.user_id = $1 and lists.project_id = $2"
	GetAllUsersColumnsByProject = "select lists.id, lists.cname, lists.index, lists.created_at from projects left join lists on lists.project_id = projects.id where projects.user_id = $1 and lists.project_id = $2"
	GetColumnById               = "select lists.id, lists.cname, lists.index, lists.created_at from projects left join lists on lists.project_id = projects.id where projects.user_id = $1 and lists.project_id = $2 and lists.id = $3"
	GetColumnIndexById          = "select lists.index from projects left join lists on lists.project_id = projects.id where projects.user_id = $1 and lists.project_id = $2 and lists.id = $3"
	IncrementColumnIndexes      = "update lists set index = index + 1 from projects where lists.project_id = projects.id and projects.user_id = $1 and lists.project_id = $2 and lists.index >= $3 and lists.index < $4"
	DecrementColumnIndexes      = "update lists set index = index - 1 from projects where lists.project_id = projects.id and projects.user_id = $1 and lists.project_id = $2 and lists.index <= $3 and lists.index > $4"
	UpdateColumnIndex           = "update lists set index = $4 from projects where lists.project_id = projects.id and projects.user_id = $1 and lists.project_id = $2 and lists.id = $3"
	DeleteColumnById            = `DELETE FROM lists using projects WHERE lists.project_id = projects.id and projects.user_id = $1 and lists.project_id = $2 and lists.id = $3`
)

// Database variables
var connString string

func NewDb(postgres config.Postgres) {
	//user=test password=test dbname=test sslmode=disable
	connString = fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", postgres.Host, postgres.Port, postgres.Database, postgres.User, postgres.Password)

	logger.Get().Info(connString)
	GetConn()
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
