package models

import "time"

type Task struct {
	Id          string    `json:"id"`
	ColumnId    string    `json:"columnId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Created     time.Time `json:"created"`
}
