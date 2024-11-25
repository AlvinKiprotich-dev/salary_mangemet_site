package models

type WorkLog struct {
	ID        string  `json:"id,omitempty" bson:"id,omitempty"`
	UserID    string  `json:"user_id,omitempty" bson:"user_id,omitempty"`
	HoursWorked float64 `json:"hours_worked,omitempty" bson:"hours_worked,omitempty"`
	Date      string  `json:"date,omitempty" bson:"date,omitempty"`
}
