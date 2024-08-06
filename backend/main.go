package main

import (
	"github.com/A12ise/SA_67_Order/tree/main/week5/backend/entity"
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
	&entity.StatusOrder{}, 
	&entity.Order{}, 
	&entity.Product{}, 
	&entity.OrderProduct{},
)

}