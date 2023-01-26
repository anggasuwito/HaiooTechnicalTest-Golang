package repository

import (
	"gorm.io/gorm"
)

type CartRepository struct {
	DB *gorm.DB
}

type CartRepositoryI interface {
}

func NewCartRepository(db *gorm.DB) CartRepositoryI {
	return &CartRepository{
		DB: db,
	}
}
