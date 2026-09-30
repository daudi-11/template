package database

import (
	"fmt"

	"test/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	//DSN stand for data source name that carries the details for db
	dsn := "host=localhost user=postgres password=12345 dbname=postgres port=5432 sslmode=disable TimeZone=UTC"

	//connection part
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to access database connection: %w", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	//this is the automigration part of the code which will create the table in the database if it does not exist
	err = db.AutoMigrate(&models.Parent{}, &models.Child{}, &models.Wallet{}, &models.Transaction{}, &models.Expense{})
	if err != nil {
		fmt.Println("Failed to run migrations:", err)
		return nil, err
	}

	return db, nil
}
