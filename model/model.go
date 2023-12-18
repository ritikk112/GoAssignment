package model

type Employee struct {
	ID   int    `json:"ID"`
	Name string `json:"name" validate:"required"`
	Dept string `json:"dept" validate:"required"`
}

