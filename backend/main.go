package main

import (
	"github.com/A12ise/SA_67_Order/entity"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)


func main() {
  db, err := gorm.Open(sqlite.Open("sa_order.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(
	&entity.Booking{}, 
	&entity.Employee{}, 
	&entity.Status_Order{}, 
	&entity.Order{}, 
	&entity.Product{}, 
	&entity.Order_Product{},
)

}