package main

import (
	"myapp/model"

	"myapp/database"

)

func main() {

	// Migrate the schema
	database.DB.AutoMigrate(&model.Todo{})

	// Create
	database.DB.Create(&model.Todo{First_name: "F-Bot01", Last_name: "L-Bot01", Task: "Prorgramer"})
	database.DB.Create(&model.Todo{First_name: "F-Bot02", Last_name: "L-Bot02", Task: "Devaloper"})

	

}
