package entity

import "gorm.io/gorm"


type Order struct {
	gorm.Model
	Order_id		string			`json:"order_id"`

	Booking_id 		*string


	Order_Product []Order_Product `gorm:"foreignKey:order_id"`
}