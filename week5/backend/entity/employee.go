package entity

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	Employee_id 	string		

	First_Name 		string		

	Last_Name 		string		

	Gender_id 		string		

	Username 		string		

	Password 		string		

	Position_id 	string		

	Register_date 	time.Time

	Order []Order `gorm:"foreignKey:Employee_id"`
}