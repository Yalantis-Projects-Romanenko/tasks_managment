package models

import "time"

type Project struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId,omitempty"`
	Name        string    `json:"name" validate:"required,gte=3"`
	Description string    `json:"description" validate:"required,gte=5"`
	Created     time.Time `json:"created"`
}
