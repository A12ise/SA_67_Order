package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Employee_id 	string		`json:"employee_id"`

	First_Name 		string		`json:"first_name"`

	Last_Name 		string		`json:"last_name"`

	Gender_id 		string		`json:"gender_id"`

	Username 		string		`json:"username"`

	Password 		string		`json:"password"`

	Position_id 	string		`json:"position_id"`

	Register_date 	time.Time	`json:"register_date"`

	Order []Order `gorm:"foreignKey:employee_id"`
}