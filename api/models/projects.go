package models

import "time"

type Project struct {
	Id          string    `json:"id"`
	UserId      string    `json:"userId"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Created     time.Time `json:"created"`
}
