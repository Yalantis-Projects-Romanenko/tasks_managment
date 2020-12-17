package models

import "time"

type Project struct {
	id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	UserId      string    `json:"userId"`
	Created     time.Time `json:"created"`
}
