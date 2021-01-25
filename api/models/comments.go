package models

import "time"

type Comment struct {
	Id        string    `json:"id"`
	ProjectId string    `json:"projectId,omitempty"`
	TaskId    string    `json:"taskId,omitempty"`
	UserId    string    `json:"username"`
	Content   string    `json:"content" validate:"required,gte=3"`
	Created   time.Time `json:"created"`
}
