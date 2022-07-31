package main

import (
	"myapp/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root@tcp(127.0.0.1:3306)/go_orm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.Todo{})

	// Create
	db.Create(&model.Todo{Fname: "F-Bot01", Lname: "L-Bot01", Task: "Prorgramer"})
	db.Create(&model.Todo{Fname: "F-Bot02", Lname: "L-Bot02", Task: "Devaloper"})

	//   // Read
	//   var product Employee
	//   db.First(&product, 1) // find product with integer primary key
	//   db.First(&product, "code = ?", "D42") // find product with code D42

	//   // Update - update product's price to 200
	//   db.Model(&product).Update("Price", 200)
	//   // Update - update multiple fields
	//   db.Model(&product).Updates(Employee{Price: 200, Code: "F42"}) // non-zero fields
	//   db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//   // Delete - delete product
	//   db.Delete(&product, 1)
}
