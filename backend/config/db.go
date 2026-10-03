package config

import (
	"fmt"

	"holygrail-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Auth *gorm.DB
}

func Connect(settings Settings) (*Database, error) {
	// [STEP 1.3] AuthDB connection handle for authentication data.
	authDB, err := openPostgres(settings.AuthDBURL)
	if err != nil {
		return nil, fmt.Errorf("connect auth database: %w", err)
	}

	if err := authDB.AutoMigrate(&models.AdminUser{}, &models.AuthCredential{}, &models.SessionToken{}); err != nil {
		closeDB(authDB)
		return nil, fmt.Errorf("migrate authentication database: %w", err)
	}

	return &Database{Auth: authDB}, nil
}

func openPostgres(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		closeDB(db)
		return nil, err
	}
	return db, nil
}

func (database *Database) Close() {
	closeDB(database.Auth)
}

func closeDB(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err == nil {
		_ = sqlDB.Close()
	}
}
