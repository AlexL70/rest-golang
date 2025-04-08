package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDB(connStr string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Post{})

	return db, nil
}
