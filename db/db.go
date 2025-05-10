package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "user=postgres password=jeetu@9712 dbname=echoDB sslmode=disable port=5433"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	fmt.Println("Connected to PostgreSQL using GORM!")
	// Automatically create tables based on your structs
	err = DB.AutoMigrate(&User{}, &Message{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
}
