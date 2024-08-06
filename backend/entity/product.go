package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_id   string 

	Product_name string	

	Category_id  string	

	Employee_id  string	

	Order_Product []Order_Product `gorm:"foreignKey:Product_id"`

}
