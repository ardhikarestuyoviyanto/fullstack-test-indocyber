package db

import (
	"auth_test/config"
	"auth_test/internal/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	conf := config.InitConfiguration()

	// Format DSN PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		conf.DB_HOST,
		conf.DB_USERNAME,
		conf.DB_PASSWORD,
		conf.DB_NAME,
		conf.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Auto migrate your models
	if err := db.AutoMigrate(&entity.Users{}); err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	log.Println("Database connected ✅")
	DB = db
	return db
}
