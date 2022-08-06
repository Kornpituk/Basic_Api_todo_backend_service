package main

import (
	"myapp/model"

	"myapp/database"
)

func main() {

	// Migrate the schema
	database.DB.AutoMigrate(&model.Todo{})

	// Create
	database.DB.Create(&model.Todo{Fname: "F-Bot01", Lname: "L-Bot01", Task: "Prorgramer"})
	database.DB.Create(&model.Todo{Fname: "F-Bot02", Lname: "L-Bot02", Task: "Devaloper"})

}
