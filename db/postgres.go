package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"recruitment-system/models"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Profile{}, &models.Job{}, &models.Application{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
