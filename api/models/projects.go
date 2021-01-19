package models

import "time"

type Project struct {
	Id string `json:"id"`
	//UserId      string    `json:"userId"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	Created     time.Time `json:"created"`
}
