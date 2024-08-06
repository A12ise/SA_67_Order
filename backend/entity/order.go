package entity

import "gorm.io/gorm"


type Order struct {
	gorm.Model
	Order_id		string		

	Booking_id 		*string
	Booking Booking `gorm:"foreignKey:Booking_id"`

	Employee_id 	*string
	Employee Employee `gorm:"foreignKey:Employee_id"`

	Status_id 		*string
	Status_Order Status_Order `gorm:"foreignKey:Status_id"`


	Order_Product []Order_Product `gorm:"foreignKey:Order_id"`
}