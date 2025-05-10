package db

import "time"

type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique;not null"`
	PasswordHash string `gorm:"column:password_hash;not null"`
}

type Message struct {
	ID         uint `gorm:"primaryKey"`
	SenderID   uint
	ReceiverID uint
	Message    string    `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
}
