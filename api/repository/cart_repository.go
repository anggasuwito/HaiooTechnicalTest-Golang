package repository

import (
	"HaiooTechnicalTest-Golang/models"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type CartRepository struct {
	DB *gorm.DB
}

type CartRepositoryI interface {
	CreateProduct(data models.Cart) error
	DeleteProduct(kodeProduk string) error
	GetProduct(param models.CartRequestParameter) ([]models.Cart, error)
}

func NewCartRepository(db *gorm.DB) CartRepositoryI {
	return &CartRepository{
		DB: db,
	}
}

func (r CartRepository) CreateProduct(data models.Cart) error {
	return r.DB.Debug().Create(&data).Error
}

func (r CartRepository) DeleteProduct(kodeProduk string) error {
	return r.DB.Debug().Where("kode_produk = ?", kodeProduk).Delete(&models.Cart{}).Error
}

func (r CartRepository) GetProduct(param models.CartRequestParameter) ([]models.Cart, error) {
	var res []models.Cart
	var where []string
	if param.NamaProduk != "" {
		where = append(where, fmt.Sprintf("nama_produk = '%v'", param.NamaProduk))
	}
	if param.Kuantitas != 0 {
		where = append(where, fmt.Sprintf("kuantitas = %v", param.Kuantitas))
	}
	err := r.DB.Debug().Where(strings.Join(where, " AND ")).Find(&res).Error
	return res, err
}
