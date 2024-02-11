package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dsn := os.Getenv("DSN")

	// This is main command to connect to mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect to the database ")
	} else {
		log.Println("Connect successfully")
	}

	DB = db
}
