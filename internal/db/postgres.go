package db

import (
	"fmt"
	"movie-reservation-system/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg models.Config) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&models.User{}); err != nil {
		panic(err)
	}

	fmt.Printf("db migrated\n")

	DB = db
}
