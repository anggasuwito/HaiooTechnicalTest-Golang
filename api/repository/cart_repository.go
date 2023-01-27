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
	FindByCode(kodeProduk string) (models.Cart, error)
	FindByName(namaProduk string) (models.Cart, error)
	UpdateProduct(data models.Cart) error
}

func NewCartRepository(db *gorm.DB) CartRepositoryI {
	return &CartRepository{
		DB: db,
	}
}

func (r CartRepository) FindByName(namaProduk string) (models.Cart, error) {
	var res models.Cart
	err := r.DB.Debug().Model(&models.Cart{}).Where("nama_produk = ?", namaProduk).Scan(&res).Error
	return res, err
}

func (r CartRepository) FindByCode(kodeProduk string) (models.Cart, error) {
	var res models.Cart
	err := r.DB.Debug().Model(&models.Cart{}).Where("kode_produk = ?", kodeProduk).Scan(&res).Error
	return res, err
}

func (r CartRepository) UpdateProduct(data models.Cart) error {
	return r.DB.Debug().Model(&models.Cart{}).Where("kode_produk = ?", data.KodeProduk).Updates(map[string]interface{}{
		"kuantitas": data.Kuantitas,
	}).Error
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
	err := r.DB.Debug().Model(&models.Cart{}).Select("*").Where(strings.Join(where, " AND ")).Scan(&res).Error
	return res, err
}
