package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_id   string `json:"product_id"`

	Product_name string	`json:"product_name"`

	Category_id  string	`json:"category_id"`

	Employee_id  string	`json:"employee_id"`

	Order_Product []Order_Product `gorm:"foreignKey:product_id"`

}
