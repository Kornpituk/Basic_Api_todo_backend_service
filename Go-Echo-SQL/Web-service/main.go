package main

import (
	"log"
	"myapp/database"
	"myapp/routes"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load("C:/xampp/htdocs/code/Go-Echo-SQL/Web-service/etc/env")

	if err != nil {
		log.Printf("%s", err)
	}

	database.ConnectDB()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1323"))
}
