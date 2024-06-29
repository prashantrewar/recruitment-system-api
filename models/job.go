package models

import "time"

type Job struct {
	ID               uint      `gorm:"primaryKey"`
	Title            string    `gorm:"not null"`
	Description      string    `gorm:"not null"`
	PostedOn         time.Time `gorm:"not null"`
	TotalApplications int
	CompanyName      string    `gorm:"not null"`
	PostedBy         uint      `gorm:"not null"` // Foreign Key to User
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
