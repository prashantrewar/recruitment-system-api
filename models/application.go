package models

import "time"

type Application struct {
	ID         uint      `gorm:"primaryKey"`
	JobID      uint      `gorm:"not null"`
	ApplicantID uint      `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
