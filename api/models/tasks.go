package models

import "time"

type Task struct {
	Id          string    `json:"id"`
	ProjectId   string    `json:"projectId,omitempty"`
	ColumnId    string    `json:"columnId,omitempty"`
	Title       string    `json:"title" validate:"required,gte=3"`
	Description string    `json:"description" validate:"required,gte=3"`
	Priority    *int      `json:"priority"`
	Created     time.Time `json:"created"`
}
