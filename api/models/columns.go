package models

import "time"

type Column struct {
	Id        string    `json:"id"`
	ProjectId string    `json:"projectId,omitempty"`
	Name      string    `json:"name" validate:"required,gte=3"`
	Index     *int64    `json:"index"` // position in project
	Created   time.Time `json:"created"`
}
