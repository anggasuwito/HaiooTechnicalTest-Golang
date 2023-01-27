package config

import (
	"HaiooTechnicalTest-Golang/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func GormDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=Asia/Jakarta",
		os.Getenv("PG_HOST"),
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_DBNAME"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_SSLMODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Cart{})
	return db
}
