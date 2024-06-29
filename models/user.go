package models

import "time"

type User struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `gorm:"not null"`
	Email           string    `gorm:"unique;not null"`
	Address         string    `gorm:"not null"`
	UserType        string    `gorm:"not null"` // Admin or Applicant
	PasswordHash    string    `gorm:"not null"`
	ProfileHeadline string
	Profile         Profile   `gorm:"foreignKey:UserID"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
