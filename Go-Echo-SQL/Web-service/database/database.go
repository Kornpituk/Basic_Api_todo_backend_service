package database

import (
	// "os"

	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	
)

var DB *gorm.DB
var err error
	
func ConnectDB() {
	dsn := os.Getenv("DNS")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
