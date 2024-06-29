package models

import "time"

type Profile struct {
	ID         uint   `gorm:"primaryKey"`
	UserID     uint   `gorm:"not null"`
	ResumeFile string `gorm:"not null"`
	Name       string
	Email      string
	Phone      string
	Skills     string
	Education  string
	Experience string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
