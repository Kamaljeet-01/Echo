package model

import (
	"echo/db"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func InsertUser(username string, passwordHash string) error {
	user := db.User{
		Username:     username,
		PasswordHash: passwordHash,
	}

	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("User inserted successfully with ID:", user.ID)
	return nil
}

func Checkuser(username string, password string) (bool, error) {
	var user db.User
	result := db.DB.Where("username = ? AND password = ?", username, password).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil // User not found
		}
		return false, result.Error // DB error
	}
	return true, nil // User exists
}
