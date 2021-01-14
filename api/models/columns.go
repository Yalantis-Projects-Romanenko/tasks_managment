package models

import "time"

type Column struct {
	Id             string    `json:"id"`
	ProjectId      string    `json:"projectId"`
	Name           string    `json:"name"`
	IndexInProject string    `json:"indexInProject"`
	Created        time.Time `json:"created"`
}
