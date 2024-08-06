package entity

import "gorm.io/gorm"

type Order_Product struct {
	gorm.Model
	Order_id    int   
	Product_id  string
	Category_id string
	Quantity   int 
	Order      Order  `gorm:"foreignKey:Order_id"`
	Product    Product `gorm:"foreignKey:Product_id"`
}