package models

import "time"

type Comment struct {
	Id       string    `json:"id"`
	TaskId   string    `json:"taskId"`
	Username string    `json:"username"`
	Content  string    `json:"content"`
	Created  time.Time `json:"created"`
}
