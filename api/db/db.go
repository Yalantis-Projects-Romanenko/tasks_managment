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
	ExistsProject       = "select exists (select id from projects where user_id = $1 and id = $2 limit 1)"
	InsertProject       = "insert into projects (pname, pdescription, user_id) values($1,$2,$3) RETURNING id"
	UpdateProject       = `UPDATE projects SET pname=$3, pdescription=$4 WHERE user_id=$1 and id=$2`

	InsertColumn                   = "insert into lists (cname, index, project_id) values($1,$2,$3) RETURNING id"
	GetMaxColumnIndex              = "select MAX (index) from lists where project_id = $1"
	GetAllUsersColumnsByProject    = "select id, cname, index, created_at from lists where project_id = $1"
	GetColumnById                  = "select id, cname, index, created_at from  lists where project_id = $1 and id = $2"
	GetColumnIndexById             = "select index from lists where project_id = $1 and id = $2"
	IncrementColumnIndexes         = "update lists set index = index + 1 where project_id = $1 and index >= $2 and index < $3"
	DecrementColumnIndexes         = "update lists set index = index - 1 where project_id = $1 and index <= $2 and index > $3"
	DecrementColumnIndexesOnDelete = "update lists set index = index - 1 where project_id = $1 and index >= $2"
	UpdateColumnIndex              = "update lists set index = $3 where project_id = $1 and id = $2"
	DeleteColumnById               = `delete from lists where project_id = $1 and id = $2 returning index`
	UpdateColumnName               = `UPDATE lists SET cname = $3 WHERE project_id = $1 and id = $2`

	GetMaxTaskPriority               = "SELECT COALESCE((select MAX (priority) from tasks where tasks.column_id = $1 and tasks.project_id = $2),-1)"
	InsertTask                       = "insert into tasks (title, description, priority, project_id,column_id) values($1,$2,$3,$4,$5) RETURNING id"
	UpdateTaskTitle                  = `UPDATE tasks SET title=$1 WHERE id=$2 and project_id = $3`
	UpdateTaskDescription            = `UPDATE tasks SET description=$1 WHERE id=$2 and project_id = $3`
	GetTasksPriorityById             = "select priority from tasks where column_id = $1 and project_id = $2 and id = $3"
	IncrementTasksPriorities         = "update tasks set priority = priority + 1 where column_id = $1 and project_id = $2 and priority >= $3 and priority < $4"
	DecrementTasksPriorities         = "update tasks set priority = priority - 1 where column_id = $1 and project_id = $2 and priority <= $3 and priority > $4"
	DecrementTasksPrioritiesOnDelete = "update tasks set priority = priority - 1 where project_id = $1 and column_id = $2 and priority >= $3"
	UpdateTasksPriority              = "update tasks set priority = $4 where column_id = $1 and  project_id = $2 and id = $3"
	GetAllTasks                      = "select id, title, description, priority, created_at from tasks where project_id = $1 and column_id = $2"
	DeleteTaskById                   = `DELETE FROM tasks WHERE project_id = $1 and column_id = $2 and id = $3 returning priority`
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
