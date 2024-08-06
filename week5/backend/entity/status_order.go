package entity

import "gorm.io/gorm"

type Status_Order struct {
	gorm.Model
	Status_id int `json:"status_id"`

	Status_name string `json:"status_name"`

	Order []Order `gorm:"foreignKey:status_id"`
}
