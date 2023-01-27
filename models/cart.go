package models

import "time"

func (Cart) TableName() string {
	return "cart"
}

type Cart struct {
	KodeProduk string    `json:"kode_produk" gorm:"type:uuid;not null;primaryKey"`
	NamaProduk string    `json:"nama_produk"`
	Kuantitas  int       `json:"kuantitas"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:current_timestamp;not null"`
}

type CartRequestBody struct {
	NamaProduk string `json:"nama_produk" binding:"required"`
	Kuantitas  int    `json:"kuantitas" binding:"required"`
}

type CartRequestParameter struct {
	NamaProduk string `query:"nama_produk"`
	Kuantitas  int    `query:"kuantitas"`
}
