package models

type User struct {
	ID     string  `json:"id,omitempty" bson:"id,omitempty"`
	Name   string  `json:"name,omitempty" bson:"name,omitempty"`
	Email  string  `json:"email,omitempty" bson:"email,omitempty"`
	Salary float64 `json:"salary,omitempty" bson:"salary,omitempty"`
}
