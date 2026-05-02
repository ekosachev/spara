package database

import (
	"fmt"

	"github.com/ekosachev/spara/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDb() (*gorm.DB, error) {
	config := config.GetConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		return nil, err
	}

	db.AutoMigrate()

	return db, nil
}
