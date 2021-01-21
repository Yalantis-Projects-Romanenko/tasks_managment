package models

type ColumnMove struct {
	Id    string `json:"id"  validate:"required"`
	Index int64  `json:"index" validate:"required,gte=0"` // position in project
}
