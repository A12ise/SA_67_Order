package entity

import (
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Booking_id 			string		`json:"booking_id"`
	Number_of_customers int
	Date 				time.Time
	Soup_id_1 			string
	Soup_id_2 			string
	Package 			string
	Table_id 			string
	Member_id 			string
	Employee_id 		string

	Order []Order `gorm:"foreignKey:booking_id"`
}